/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package initializer_test

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"pkg.deepin.io/lib/initializer/v2"
	"testing"
)

func e1() error {
	return nil
}

func e2() error {
	var err error
	return err
}

func TestInitializer(t *testing.T) {
	Convey("test initializer with success", t, func() {
		err := initializer.Do(func() error {
			return nil
		}).Do(func() error {
			return e1()
		}).Do(func() error {
			return e2()
		}).GetError()

		So(err, ShouldBeNil)
	})
}

func TestInitializerError(t *testing.T) {
	Convey("test initializer with error", t, func() {
		var err error
		So(func() {
			err = initializer.Do(func() error {
				return e1()
			}).Do(func() error {
				return e2()
			}).Do(func() error {
				return errors.New("initialize error")
			}).Do(func() error {
				panic("should not be executed")
				return nil
			}).GetError()

		}, ShouldNotPanic)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "initialize error")
	})
}
