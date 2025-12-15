/*
 * Copyright (c) 2025 KAnggara
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara on Sunday 14/12/2025 23.03
 * @project PakaiWA
 * ~/work/PakaiWA/PakaiWA/internal/app/pakaiwa/error
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/app/pakaiwa/error
 */

package apperror

import (
	"errors"
	"fmt"
)

func PanicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

var ErrUsernameExists = errors.New("username already exists")
var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrPasswordWeak = errors.New("password does not meet complexity requirements")
