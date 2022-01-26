package util

import (
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
)

// Win1251ToUTF8 - Русские фразы из базы нужно конвертировать из win1251 в utf8 этим:
func Win1251ToUTF8(str string) (string, error) {
	tr := transform.NewReader(strings.NewReader(str), charmap.Windows1251.NewDecoder())
	buf, err := ioutil.ReadAll(tr)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

/* Количество кадров стека, которые необходимо пропустить перед записью на ПК, где 0 идентифицирует
кадр для самих вызывающих абонентов, а 1 идентифицирует вызывающего абонента. Возвращает количество
записей, записанных на компьютер.*/
const skipNumOfStackFrame = 3

// FileWithLineNum Возвращает имя файла и номер строки текущего файла.
func FileWithLineNum() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skipNumOfStackFrame, pc)
	frame, _ := runtime.CallersFrames(pc[:n]).Next()
	idxFile := strings.LastIndexByte(frame.File, '/')

	return frame.File[idxFile+1:] + ":" + strconv.FormatInt(int64(frame.Line), 10)
}

// FileWithFuncAndLineNum Возвращает имя файла с функцией и номер строки текущего файла.
func FileWithFuncAndLineNum() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skipNumOfStackFrame, pc)
	frame, _ := runtime.CallersFrames(pc[:n]).Next()
	idxFile := strings.LastIndexByte(strconv.Itoa(frame.Line), '/')

	return "[" + frame.Function + "] - " + frame.File[idxFile+1:] + ":" + strconv.FormatInt(int64(frame.Line), 10)
}
