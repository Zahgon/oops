package oops

///
/// Generic wrapper functions for handling multiple return values with error wrapping.
///
/// This module provides a set of generic functions that simplify error handling
/// when functions return multiple values along with an error. These functions
/// automatically wrap the error using oops.Wrap or oops.Wrapf while preserving
/// all other return values.
///
/// The functions follow a naming convention where the number indicates the
/// total number of return values (including the error). For example, Wrap2
/// handles functions that return 2 values (1 value + 1 error), Wrap3 handles
/// 3 values (2 values + 1 error), and so on up to Wrap10.
///
/// These functions are particularly useful for:
/// - Reducing boilerplate code in error handling
/// - Ensuring consistent error wrapping across the codebase
/// - Simplifying function calls that return multiple values
/// - Maintaining clean, readable code when dealing with complex return types
///

// Wrap2 wraps an error while preserving a single return value.
// This function is useful for functions that return (value, error) pairs
// and need the error to be wrapped with oops error handling.
//
// The function takes a value of any type A and an error, wraps the error
// using oops.Wrap, and returns the original value along with the wrapped error.
//
// Example usage:
//
//	func getUser(id string) (User, error) {
//	    user, err := database.GetUser(id)
//	    return oops.Wrap2(user, err)
//	}
//
//	// Equivalent to:
//	func getUser(id string) (User, error) {
//	    user, err := database.GetUser(id)
//	    if err != nil {
//	        return user, oops.Wrap(err)
//	    }
//	    return user, nil
//	}
func Wrap2[A any](a A, err error) (A, error) {
	_ = "STUB: not implemented"
	return *

	// Wrap3 wraps an error while preserving two return values.
	// This function is useful for functions that return (value1, value2, error)
	// and need the error to be wrapped with oops error handling.
	//
	// The function takes two values of any types A and B and an error, wraps
	// the error using oops.Wrap, and returns the original values along with
	// the wrapped error.
	//
	// Example usage:
	//
	//	func getUserAndProfile(id string) (User, Profile, error) {
	//	    user, profile, err := database.GetUserAndProfile(id)
	//	    return oops.Wrap3(user, profile, err)
	//	}
	new(A), nil
}

func Wrap3[A any, B any](a A, b B, err error) (A, B, error) {
	_ = "STUB: not implemented"
	return *

	// Wrap4 wraps an error while preserving three return values.
	// This function is useful for functions that return (value1, value2, value3, error)
	// and need the error to be wrapped with oops error handling.
	//
	// The function takes three values of any types A, B, and C and an error,
	// wraps the error using oops.Wrap, and returns the original values along
	// with the wrapped error.
	//
	// Example usage:
	//
	//	func getUserProfileAndSettings(id string) (User, Profile, Settings, error) {
	//	    user, profile, settings, err := database.GetUserProfileAndSettings(id)
	//	    return oops.Wrap4(user, profile, settings, err)
	//	}
	new(A), *new(B), nil
}

func Wrap4[A any, B any, C any](a A, b B, c C, err error) (A, B, C, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), nil
}

// Wrap5 wraps an error while preserving four return values.
// This function is useful for functions that return (value1, value2, value3, value4, error)
// and need the error to be wrapped with oops error handling.
//
// The function takes four values of any types A, B, C, and D and an error,
// wraps the error using oops.Wrap, and returns the original values along
// with the wrapped error.
func Wrap5[A any, B any, C any, D any](a A, b B, c C, d D, err error) (A, B, C, D, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), nil
}

// Wrap6 wraps an error while preserving five return values.
// This function is useful for functions that return (value1, value2, value3, value4, value5, error)
// and need the error to be wrapped with oops error handling.
//
// The function takes five values of any types A, B, C, D, and E and an error,
// wraps the error using oops.Wrap, and returns the original values along
// with the wrapped error.
func Wrap6[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E, err error) (A, B, C, D, E, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), nil
}

// Wrap7 wraps an error while preserving six return values.
// This function is useful for functions that return (value1, value2, value3, value4, value5, value6, error)
// and need the error to be wrapped with oops error handling.
//
// The function takes six values of any types A, B, C, D, E, and F and an error,
// wraps the error using oops.Wrap, and returns the original values along
// with the wrapped error.
func Wrap7[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F, err error) (A, B, C, D, E, F, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), nil
}

// Wrap8 wraps an error while preserving seven return values.
// This function is useful for functions that return (value1, value2, value3, value4, value5, value6, value7, error)
// and need the error to be wrapped with oops error handling.
//
// The function takes seven values of any types A, B, C, D, E, F, and G and an error,
// wraps the error using oops.Wrap, and returns the original values along
// with the wrapped error.
func Wrap8[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G, err error) (A, B, C, D, E, F, G, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), nil
}

// Wrap9 wraps an error while preserving eight return values.
// This function is useful for functions that return (value1, value2, value3, value4, value5, value6, value7, value8, error)
// and need the error to be wrapped with oops error handling.
//
// The function takes eight values of any types A, B, C, D, E, F, G, and H and an error,
// wraps the error using oops.Wrap, and returns the original values along
// with the wrapped error.
func Wrap9[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H, err error) (A, B, C, D, E, F, G, H, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), *new(H), nil
}

// Wrap10 wraps an error while preserving nine return values.
// This function is useful for functions that return (value1, value2, value3, value4, value5, value6, value7, value8, value9, error)
// and need the error to be wrapped with oops error handling.
//
// The function takes nine values of any types A, B, C, D, E, F, G, H, and I and an error,
// wraps the error using oops.Wrap, and returns the original values along
// with the wrapped error.
func Wrap10[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I, err error) (A, B, C, D, E, F, G, H, I, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), *new(H), *new(I), nil
}

// Wrapf2 wraps an error with a formatted message while preserving a single return value.
// This function is similar to Wrap2 but uses oops.Wrapf to add additional context
// to the error message.
//
// The function takes a value of any type A, an error, and a format string with
// arguments, wraps the error using oops.Wrapf, and returns the original value
// along with the wrapped error.
//
// Example usage:
//
//	func getUser(id string) (User, error) {
//	    user, err := database.GetUser(id)
//	    return oops.Wrapf2(user, err, "failed to get user with id %s", id)
//	}
func Wrapf2[A any](a A, err error, format string, args ...any) (A, error) {
	_ = "STUB: not implemented"
	return *new(A), nil
}

// Wrapf3 wraps an error with a formatted message while preserving two return values.
// This function is similar to Wrap3 but uses oops.Wrapf to add additional context
// to the error message.
//
// The function takes two values of any types A and B, an error, and a format string
// with arguments, wraps the error using oops.Wrapf, and returns the original values
// along with the wrapped error.
//
// Example usage:
//
//	func getUserAndProfile(id string) (User, Profile, error) {
//	    user, profile, err := database.GetUserAndProfile(id)
//	    return oops.Wrapf3(user, profile, err, "failed to get user and profile for id %s", id)
//	}
func Wrapf3[A any, B any](a A, b B, err error, format string, args ...any) (A, B, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), nil
}

// Wrapf4 wraps an error with a formatted message while preserving three return values.
// This function is similar to Wrap4 but uses oops.Wrapf to add additional context
// to the error message.
//
// The function takes three values of any types A, B, and C, an error, and a format
// string with arguments, wraps the error using oops.Wrapf, and returns the original
// values along with the wrapped error.
func Wrapf4[A any, B any, C any](a A, b B, c C, err error, format string, args ...any) (A, B, C, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), nil
}

// Wrapf5 wraps an error with a formatted message while preserving four return values.
// This function is similar to Wrap5 but uses oops.Wrapf to add additional context
// to the error message.
//
// The function takes four values of any types A, B, C, and D, an error, and a format
// string with arguments, wraps the error using oops.Wrapf, and returns the original
// values along with the wrapped error.
func Wrapf5[A any, B any, C any, D any](a A, b B, c C, d D, err error, format string, args ...any) (A, B, C, D, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), nil
}

// Wrapf6 wraps an error with a formatted message while preserving five return values.
// This function is similar to Wrap6 but uses oops.Wrapf to add additional context
// to the error message.
//
// The function takes five values of any types A, B, C, D, and E, an error, and a format
// string with arguments, wraps the error using oops.Wrapf, and returns the original
// values along with the wrapped error.
func Wrapf6[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E, err error, format string, args ...any) (A, B, C, D, E, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), nil
}

// Wrapf7 wraps an error with a formatted message while preserving six return values.
// This function is similar to Wrap7 but uses oops.Wrapf to add additional context
// to the error message.
//
// The function takes six values of any types A, B, C, D, E, and F, an error, and a format
// string with arguments, wraps the error using oops.Wrapf, and returns the original
// values along with the wrapped error.
func Wrapf7[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F, err error, format string, args ...any) (A, B, C, D, E, F, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), nil
}

// Wrapf8 wraps an error with a formatted message while preserving seven return values.
// This function is similar to Wrap8 but uses oops.Wrapf to add additional context
// to the error message.
//
// The function takes seven values of any types A, B, C, D, E, F, and G, an error, and a format
// string with arguments, wraps the error using oops.Wrapf, and returns the original
// values along with the wrapped error.
func Wrapf8[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G, err error, format string, args ...any) (A, B, C, D, E, F, G, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), nil
}

// Wrapf9 wraps an error with a formatted message while preserving eight return values.
// This function is similar to Wrap9 but uses oops.Wrapf to add additional context
// to the error message.
//
// The function takes eight values of any types A, B, C, D, E, F, G, and H, an error, and a format
// string with arguments, wraps the error using oops.Wrapf, and returns the original
// values along with the wrapped error.
func Wrapf9[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H, err error, format string, args ...any) (A, B, C, D, E, F, G, H, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), *new(H), nil
}

// Wrapf10 wraps an error with a formatted message while preserving nine return values.
// This function is similar to Wrap10 but uses oops.Wrapf to add additional context
// to the error message.
//
// The function takes nine values of any types A, B, C, D, E, F, G, H, and I, an error, and a format
// string with arguments, wraps the error using oops.Wrapf, and returns the original
// values along with the wrapped error.
func Wrapf10[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I, err error, format string, args ...any) (A, B, C, D, E, F, G, H, I, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), *new(C), *new(D), *new(E), *new(F), *new(G), *new(H), *new(I), nil
}
