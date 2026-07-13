package errors

import "errors"

var ErrProjectNotFound = errors.New("project not found")

// Permite posteriormente fazer if errors.Is(err, errors.ErrProjectNotFound)
// para verificar se o erro é do tipo ErrProjectNotFound.
