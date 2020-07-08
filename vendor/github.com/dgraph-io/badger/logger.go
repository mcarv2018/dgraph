/*
 * Copyright 2018 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package badger

import (
	"log"
	"os"
)

// Logger is implemented by any logging system that is used for standard logs.
type Logger interface {
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
	Warningf(string, ...interface{})
}

var badgerLogger Logger

// SetLogger sets the global logger.
func SetLogger(l Logger) { badgerLogger = l }

// Errorf logs an ERROR message to the global logger.
func Errorf(format string, v ...interface{}) {
	badgerLogger.Errorf(format, v...)
}

// Errorf logs an ERROR log message to the logger specified in opts or to the
// global logger if no logger is specified in opts.
func (opt *Options) Errorf(format string, v ...interface{}) {
	if opt.Logger != nil {
		opt.Logger.Errorf(format, v...)
		return
	}
	Errorf(format, v...)
}

// Infof logs an INFO message to the global logger.
func Infof(format string, v ...interface{}) {
	badgerLogger.Infof(format, v...)
}

// Infof is like Errorf but for INFO messages.
func (opt *Options) Infof(format string, v ...interface{}) {
	if opt.Logger != nil {
		opt.Logger.Infof(format, v...)
		return
	}
	Infof(format, v...)
}

// Warningf logs a WARNING message to the global logger.
func Warningf(format string, v ...interface{}) {
	badgerLogger.Warningf(format, v...)
}

// Warningf is like Errorf but for WARNING messages.
func (opt *Options) Warningf(format string, v ...interface{}) {
	if opt.Logger != nil {
		opt.Logger.Warningf(format, v...)
		return
	}
	Warningf(format, v...)
}

type defaultLog struct {
	*log.Logger
}

var defaultLogger = &defaultLog{Logger: log.New(os.Stderr, "badger ", log.LstdFlags)}

// UseDefaultLogger sets the global logger to the default logger.
func UseDefaultLogger() { SetLogger(defaultLogger) }

func (l *defaultLog) Errorf(f string, v ...interface{}) {
	l.Printf("ERROR: "+f, v...)
}

func (l *defaultLog) Infof(f string, v ...interface{}) {
	l.Printf("INFO: "+f, v...)
}

func (l *defaultLog) Warningf(f string, v ...interface{}) {
	l.Printf("WARNING: "+f, v...)
}

func init() {
	UseDefaultLogger()
}