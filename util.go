package yaptranslator

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/magiconair/properties"
	"golang.org/x/sync/semaphore"
	"gopkg.in/cheggaaa/pb.v2"
)

const (
	defaultConcurrent = 5
)

// StringifyProperties with comments
func StringifyProperties(p *properties.Properties) string {
	var w bytes.Buffer
	p.Write(&w, properties.UTF8)
	return w.String()
}

// TranslatePropertiesFile process
func TranslatePropertiesFile(props *properties.Properties, translator Translator, lang, target string) (*properties.Properties, error) {
	tmp := properties.NewProperties()

	propsKeys := props.Keys()

	progress := pb.Full.Start(len(propsKeys))

	var wg sync.WaitGroup

	ctx := context.TODO()

	sem := semaphore.NewWeighted(defaultConcurrent)

	wLock := &sync.Mutex{}

	for _, k := range propsKeys {
		// must set wg & sem outter that async func
		wg.Add(1)
		sem.Acquire(ctx, 1)
		v := props.MustGetString(k)
		go asyncTranslateValue(ctx, tmp, k, v, lang, target, translator, progress, &wg, sem, wLock)
	}

	wg.Wait()

	progress.Finish()

	for _, k := range propsKeys {
		// must set wg & sem outter that async func
		props.SetValue(k, tmp.MustGetString(k))
	}

	return props, nil
}

func asyncTranslateValue(ctx context.Context, ps *properties.Properties, k1, v1, lang, target string, translator Translator, progress *pb.ProgressBar, wg *sync.WaitGroup, sem *semaphore.Weighted, lock *sync.Mutex) {

	defer wg.Done()
	defer progress.Increment()
	defer sem.Release(1)

	if len(v1) == 0 {
		// not translate empty string
		lock.Lock()
		ps.SetValue(k1, v1)
		lock.Unlock()
	} else {
		tv, err := translator.GetTranslation(v1, lang, target)
		lock.Lock()
		if err == nil {
			ps.SetValue(k1, tv)
		} else {
			ps.SetValue(k1, v1) // fail back use original value
			log.Printf("translate '%s' failed from %s to %s: %v", v1, lang, target, err)
		}
		lock.Unlock()
	}

}

// AddSuffixToFileName string
func AddSuffixToFileName(path, suffix string) string {
	base, filename := filepath.Split(path)
	ext := filepath.Ext(filename)
	filebasename := strings.TrimSuffix(filename, ext)
	return fmt.Sprint(base, filebasename, suffix, ext)
}
