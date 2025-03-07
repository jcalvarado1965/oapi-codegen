// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package server

import (
	"net/http"
	"sync"
)

// Ensure, that ServerInterfaceMock does implement ServerInterface.
// If this is not the case, regenerate this file with moq.
var _ ServerInterface = &ServerInterfaceMock{}

// ServerInterfaceMock is a mock implementation of ServerInterface.
//
// 	func TestSomethingThatUsesServerInterface(t *testing.T) {
//
// 		// make and configure a mocked ServerInterface
// 		mockedServerInterface := &ServerInterfaceMock{
// 			CreateResourceFunc: func(w http.ResponseWriter, r *http.Request, argument Argument)  {
// 				panic("mock out the CreateResource method")
// 			},
// 			CreateResource2Func: func(w http.ResponseWriter, r *http.Request, inlineArgument int, params CreateResource2Params)  {
// 				panic("mock out the CreateResource2 method")
// 			},
// 			GetEveryTypeOptionalFunc: func(w http.ResponseWriter, r *http.Request)  {
// 				panic("mock out the GetEveryTypeOptional method")
// 			},
// 			GetReservedKeywordFunc: func(w http.ResponseWriter, r *http.Request)  {
// 				panic("mock out the GetReservedKeyword method")
// 			},
// 			GetResponseWithReferenceFunc: func(w http.ResponseWriter, r *http.Request)  {
// 				panic("mock out the GetResponseWithReference method")
// 			},
// 			GetSimpleFunc: func(w http.ResponseWriter, r *http.Request)  {
// 				panic("mock out the GetSimple method")
// 			},
// 			GetWithArgsFunc: func(w http.ResponseWriter, r *http.Request, params GetWithArgsParams)  {
// 				panic("mock out the GetWithArgs method")
// 			},
// 			GetWithContentTypeFunc: func(w http.ResponseWriter, r *http.Request, contentType GetWithContentTypeParamsContentType)  {
// 				panic("mock out the GetWithContentType method")
// 			},
// 			GetWithReferencesFunc: func(w http.ResponseWriter, r *http.Request, globalArgument int64, argument Argument)  {
// 				panic("mock out the GetWithReferences method")
// 			},
// 			UpdateResource3Func: func(w http.ResponseWriter, r *http.Request, pFallthrough int)  {
// 				panic("mock out the UpdateResource3 method")
// 			},
// 		}
//
// 		// use mockedServerInterface in code that requires ServerInterface
// 		// and then make assertions.
//
// 	}
type ServerInterfaceMock struct {
	// CreateResourceFunc mocks the CreateResource method.
	CreateResourceFunc func(w http.ResponseWriter, r *http.Request, argument Argument)

	// CreateResource2Func mocks the CreateResource2 method.
	CreateResource2Func func(w http.ResponseWriter, r *http.Request, inlineArgument int, params CreateResource2Params)

	// GetEveryTypeOptionalFunc mocks the GetEveryTypeOptional method.
	GetEveryTypeOptionalFunc func(w http.ResponseWriter, r *http.Request)

	// GetReservedKeywordFunc mocks the GetReservedKeyword method.
	GetReservedKeywordFunc func(w http.ResponseWriter, r *http.Request)

	// GetResponseWithReferenceFunc mocks the GetResponseWithReference method.
	GetResponseWithReferenceFunc func(w http.ResponseWriter, r *http.Request)

	// GetSimpleFunc mocks the GetSimple method.
	GetSimpleFunc func(w http.ResponseWriter, r *http.Request)

	// GetWithArgsFunc mocks the GetWithArgs method.
	GetWithArgsFunc func(w http.ResponseWriter, r *http.Request, params GetWithArgsParams)

	// GetWithContentTypeFunc mocks the GetWithContentType method.
	GetWithContentTypeFunc func(w http.ResponseWriter, r *http.Request, contentType GetWithContentTypeParamsContentType)

	// GetWithReferencesFunc mocks the GetWithReferences method.
	GetWithReferencesFunc func(w http.ResponseWriter, r *http.Request, globalArgument int64, argument Argument)

	// UpdateResource3Func mocks the UpdateResource3 method.
	UpdateResource3Func func(w http.ResponseWriter, r *http.Request, pFallthrough int)

	// calls tracks calls to the methods.
	calls struct {
		// CreateResource holds details about calls to the CreateResource method.
		CreateResource []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
			// Argument is the argument argument value.
			Argument Argument
		}
		// CreateResource2 holds details about calls to the CreateResource2 method.
		CreateResource2 []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
			// InlineArgument is the inlineArgument argument value.
			InlineArgument int
			// Params is the params argument value.
			Params CreateResource2Params
		}
		// GetEveryTypeOptional holds details about calls to the GetEveryTypeOptional method.
		GetEveryTypeOptional []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
		}
		// GetReservedKeyword holds details about calls to the GetReservedKeyword method.
		GetReservedKeyword []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
		}
		// GetResponseWithReference holds details about calls to the GetResponseWithReference method.
		GetResponseWithReference []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
		}
		// GetSimple holds details about calls to the GetSimple method.
		GetSimple []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
		}
		// GetWithArgs holds details about calls to the GetWithArgs method.
		GetWithArgs []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
			// Params is the params argument value.
			Params GetWithArgsParams
		}
		// GetWithContentType holds details about calls to the GetWithContentType method.
		GetWithContentType []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
			// ContentType is the contentType argument value.
			ContentType GetWithContentTypeParamsContentType
		}
		// GetWithReferences holds details about calls to the GetWithReferences method.
		GetWithReferences []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
			// GlobalArgument is the globalArgument argument value.
			GlobalArgument int64
			// Argument is the argument argument value.
			Argument Argument
		}
		// UpdateResource3 holds details about calls to the UpdateResource3 method.
		UpdateResource3 []struct {
			// W is the w argument value.
			W http.ResponseWriter
			// R is the r argument value.
			R *http.Request
			// PFallthrough is the pFallthrough argument value.
			PFallthrough int
		}
	}
	lockCreateResource           sync.RWMutex
	lockCreateResource2          sync.RWMutex
	lockGetEveryTypeOptional     sync.RWMutex
	lockGetReservedKeyword       sync.RWMutex
	lockGetResponseWithReference sync.RWMutex
	lockGetSimple                sync.RWMutex
	lockGetWithArgs              sync.RWMutex
	lockGetWithContentType       sync.RWMutex
	lockGetWithReferences        sync.RWMutex
	lockUpdateResource3          sync.RWMutex
}

// CreateResource calls CreateResourceFunc.
func (mock *ServerInterfaceMock) CreateResource(w http.ResponseWriter, r *http.Request, argument Argument) {
	if mock.CreateResourceFunc == nil {
		panic("ServerInterfaceMock.CreateResourceFunc: method is nil but ServerInterface.CreateResource was just called")
	}
	callInfo := struct {
		W        http.ResponseWriter
		R        *http.Request
		Argument Argument
	}{
		W:        w,
		R:        r,
		Argument: argument,
	}
	mock.lockCreateResource.Lock()
	mock.calls.CreateResource = append(mock.calls.CreateResource, callInfo)
	mock.lockCreateResource.Unlock()
	mock.CreateResourceFunc(w, r, argument)
}

// CreateResourceCalls gets all the calls that were made to CreateResource.
// Check the length with:
//     len(mockedServerInterface.CreateResourceCalls())
func (mock *ServerInterfaceMock) CreateResourceCalls() []struct {
	W        http.ResponseWriter
	R        *http.Request
	Argument Argument
} {
	var calls []struct {
		W        http.ResponseWriter
		R        *http.Request
		Argument Argument
	}
	mock.lockCreateResource.RLock()
	calls = mock.calls.CreateResource
	mock.lockCreateResource.RUnlock()
	return calls
}

// CreateResource2 calls CreateResource2Func.
func (mock *ServerInterfaceMock) CreateResource2(w http.ResponseWriter, r *http.Request, inlineArgument int, params CreateResource2Params) {
	if mock.CreateResource2Func == nil {
		panic("ServerInterfaceMock.CreateResource2Func: method is nil but ServerInterface.CreateResource2 was just called")
	}
	callInfo := struct {
		W              http.ResponseWriter
		R              *http.Request
		InlineArgument int
		Params         CreateResource2Params
	}{
		W:              w,
		R:              r,
		InlineArgument: inlineArgument,
		Params:         params,
	}
	mock.lockCreateResource2.Lock()
	mock.calls.CreateResource2 = append(mock.calls.CreateResource2, callInfo)
	mock.lockCreateResource2.Unlock()
	mock.CreateResource2Func(w, r, inlineArgument, params)
}

// CreateResource2Calls gets all the calls that were made to CreateResource2.
// Check the length with:
//     len(mockedServerInterface.CreateResource2Calls())
func (mock *ServerInterfaceMock) CreateResource2Calls() []struct {
	W              http.ResponseWriter
	R              *http.Request
	InlineArgument int
	Params         CreateResource2Params
} {
	var calls []struct {
		W              http.ResponseWriter
		R              *http.Request
		InlineArgument int
		Params         CreateResource2Params
	}
	mock.lockCreateResource2.RLock()
	calls = mock.calls.CreateResource2
	mock.lockCreateResource2.RUnlock()
	return calls
}

// GetEveryTypeOptional calls GetEveryTypeOptionalFunc.
func (mock *ServerInterfaceMock) GetEveryTypeOptional(w http.ResponseWriter, r *http.Request) {
	if mock.GetEveryTypeOptionalFunc == nil {
		panic("ServerInterfaceMock.GetEveryTypeOptionalFunc: method is nil but ServerInterface.GetEveryTypeOptional was just called")
	}
	callInfo := struct {
		W http.ResponseWriter
		R *http.Request
	}{
		W: w,
		R: r,
	}
	mock.lockGetEveryTypeOptional.Lock()
	mock.calls.GetEveryTypeOptional = append(mock.calls.GetEveryTypeOptional, callInfo)
	mock.lockGetEveryTypeOptional.Unlock()
	mock.GetEveryTypeOptionalFunc(w, r)
}

// GetEveryTypeOptionalCalls gets all the calls that were made to GetEveryTypeOptional.
// Check the length with:
//     len(mockedServerInterface.GetEveryTypeOptionalCalls())
func (mock *ServerInterfaceMock) GetEveryTypeOptionalCalls() []struct {
	W http.ResponseWriter
	R *http.Request
} {
	var calls []struct {
		W http.ResponseWriter
		R *http.Request
	}
	mock.lockGetEveryTypeOptional.RLock()
	calls = mock.calls.GetEveryTypeOptional
	mock.lockGetEveryTypeOptional.RUnlock()
	return calls
}

// GetReservedKeyword calls GetReservedKeywordFunc.
func (mock *ServerInterfaceMock) GetReservedKeyword(w http.ResponseWriter, r *http.Request) {
	if mock.GetReservedKeywordFunc == nil {
		panic("ServerInterfaceMock.GetReservedKeywordFunc: method is nil but ServerInterface.GetReservedKeyword was just called")
	}
	callInfo := struct {
		W http.ResponseWriter
		R *http.Request
	}{
		W: w,
		R: r,
	}
	mock.lockGetReservedKeyword.Lock()
	mock.calls.GetReservedKeyword = append(mock.calls.GetReservedKeyword, callInfo)
	mock.lockGetReservedKeyword.Unlock()
	mock.GetReservedKeywordFunc(w, r)
}

// GetReservedKeywordCalls gets all the calls that were made to GetReservedKeyword.
// Check the length with:
//     len(mockedServerInterface.GetReservedKeywordCalls())
func (mock *ServerInterfaceMock) GetReservedKeywordCalls() []struct {
	W http.ResponseWriter
	R *http.Request
} {
	var calls []struct {
		W http.ResponseWriter
		R *http.Request
	}
	mock.lockGetReservedKeyword.RLock()
	calls = mock.calls.GetReservedKeyword
	mock.lockGetReservedKeyword.RUnlock()
	return calls
}

// GetResponseWithReference calls GetResponseWithReferenceFunc.
func (mock *ServerInterfaceMock) GetResponseWithReference(w http.ResponseWriter, r *http.Request) {
	if mock.GetResponseWithReferenceFunc == nil {
		panic("ServerInterfaceMock.GetResponseWithReferenceFunc: method is nil but ServerInterface.GetResponseWithReference was just called")
	}
	callInfo := struct {
		W http.ResponseWriter
		R *http.Request
	}{
		W: w,
		R: r,
	}
	mock.lockGetResponseWithReference.Lock()
	mock.calls.GetResponseWithReference = append(mock.calls.GetResponseWithReference, callInfo)
	mock.lockGetResponseWithReference.Unlock()
	mock.GetResponseWithReferenceFunc(w, r)
}

// GetResponseWithReferenceCalls gets all the calls that were made to GetResponseWithReference.
// Check the length with:
//     len(mockedServerInterface.GetResponseWithReferenceCalls())
func (mock *ServerInterfaceMock) GetResponseWithReferenceCalls() []struct {
	W http.ResponseWriter
	R *http.Request
} {
	var calls []struct {
		W http.ResponseWriter
		R *http.Request
	}
	mock.lockGetResponseWithReference.RLock()
	calls = mock.calls.GetResponseWithReference
	mock.lockGetResponseWithReference.RUnlock()
	return calls
}

// GetSimple calls GetSimpleFunc.
func (mock *ServerInterfaceMock) GetSimple(w http.ResponseWriter, r *http.Request) {
	if mock.GetSimpleFunc == nil {
		panic("ServerInterfaceMock.GetSimpleFunc: method is nil but ServerInterface.GetSimple was just called")
	}
	callInfo := struct {
		W http.ResponseWriter
		R *http.Request
	}{
		W: w,
		R: r,
	}
	mock.lockGetSimple.Lock()
	mock.calls.GetSimple = append(mock.calls.GetSimple, callInfo)
	mock.lockGetSimple.Unlock()
	mock.GetSimpleFunc(w, r)
}

// GetSimpleCalls gets all the calls that were made to GetSimple.
// Check the length with:
//     len(mockedServerInterface.GetSimpleCalls())
func (mock *ServerInterfaceMock) GetSimpleCalls() []struct {
	W http.ResponseWriter
	R *http.Request
} {
	var calls []struct {
		W http.ResponseWriter
		R *http.Request
	}
	mock.lockGetSimple.RLock()
	calls = mock.calls.GetSimple
	mock.lockGetSimple.RUnlock()
	return calls
}

// GetWithArgs calls GetWithArgsFunc.
func (mock *ServerInterfaceMock) GetWithArgs(w http.ResponseWriter, r *http.Request, params GetWithArgsParams) {
	if mock.GetWithArgsFunc == nil {
		panic("ServerInterfaceMock.GetWithArgsFunc: method is nil but ServerInterface.GetWithArgs was just called")
	}
	callInfo := struct {
		W      http.ResponseWriter
		R      *http.Request
		Params GetWithArgsParams
	}{
		W:      w,
		R:      r,
		Params: params,
	}
	mock.lockGetWithArgs.Lock()
	mock.calls.GetWithArgs = append(mock.calls.GetWithArgs, callInfo)
	mock.lockGetWithArgs.Unlock()
	mock.GetWithArgsFunc(w, r, params)
}

// GetWithArgsCalls gets all the calls that were made to GetWithArgs.
// Check the length with:
//     len(mockedServerInterface.GetWithArgsCalls())
func (mock *ServerInterfaceMock) GetWithArgsCalls() []struct {
	W      http.ResponseWriter
	R      *http.Request
	Params GetWithArgsParams
} {
	var calls []struct {
		W      http.ResponseWriter
		R      *http.Request
		Params GetWithArgsParams
	}
	mock.lockGetWithArgs.RLock()
	calls = mock.calls.GetWithArgs
	mock.lockGetWithArgs.RUnlock()
	return calls
}

// GetWithContentType calls GetWithContentTypeFunc.
func (mock *ServerInterfaceMock) GetWithContentType(w http.ResponseWriter, r *http.Request, contentType GetWithContentTypeParamsContentType) {
	if mock.GetWithContentTypeFunc == nil {
		panic("ServerInterfaceMock.GetWithContentTypeFunc: method is nil but ServerInterface.GetWithContentType was just called")
	}
	callInfo := struct {
		W           http.ResponseWriter
		R           *http.Request
		ContentType GetWithContentTypeParamsContentType
	}{
		W:           w,
		R:           r,
		ContentType: contentType,
	}
	mock.lockGetWithContentType.Lock()
	mock.calls.GetWithContentType = append(mock.calls.GetWithContentType, callInfo)
	mock.lockGetWithContentType.Unlock()
	mock.GetWithContentTypeFunc(w, r, contentType)
}

// GetWithContentTypeCalls gets all the calls that were made to GetWithContentType.
// Check the length with:
//     len(mockedServerInterface.GetWithContentTypeCalls())
func (mock *ServerInterfaceMock) GetWithContentTypeCalls() []struct {
	W           http.ResponseWriter
	R           *http.Request
	ContentType GetWithContentTypeParamsContentType
} {
	var calls []struct {
		W           http.ResponseWriter
		R           *http.Request
		ContentType GetWithContentTypeParamsContentType
	}
	mock.lockGetWithContentType.RLock()
	calls = mock.calls.GetWithContentType
	mock.lockGetWithContentType.RUnlock()
	return calls
}

// GetWithReferences calls GetWithReferencesFunc.
func (mock *ServerInterfaceMock) GetWithReferences(w http.ResponseWriter, r *http.Request, globalArgument int64, argument Argument) {
	if mock.GetWithReferencesFunc == nil {
		panic("ServerInterfaceMock.GetWithReferencesFunc: method is nil but ServerInterface.GetWithReferences was just called")
	}
	callInfo := struct {
		W              http.ResponseWriter
		R              *http.Request
		GlobalArgument int64
		Argument       Argument
	}{
		W:              w,
		R:              r,
		GlobalArgument: globalArgument,
		Argument:       argument,
	}
	mock.lockGetWithReferences.Lock()
	mock.calls.GetWithReferences = append(mock.calls.GetWithReferences, callInfo)
	mock.lockGetWithReferences.Unlock()
	mock.GetWithReferencesFunc(w, r, globalArgument, argument)
}

// GetWithReferencesCalls gets all the calls that were made to GetWithReferences.
// Check the length with:
//     len(mockedServerInterface.GetWithReferencesCalls())
func (mock *ServerInterfaceMock) GetWithReferencesCalls() []struct {
	W              http.ResponseWriter
	R              *http.Request
	GlobalArgument int64
	Argument       Argument
} {
	var calls []struct {
		W              http.ResponseWriter
		R              *http.Request
		GlobalArgument int64
		Argument       Argument
	}
	mock.lockGetWithReferences.RLock()
	calls = mock.calls.GetWithReferences
	mock.lockGetWithReferences.RUnlock()
	return calls
}

// UpdateResource3 calls UpdateResource3Func.
func (mock *ServerInterfaceMock) UpdateResource3(w http.ResponseWriter, r *http.Request, pFallthrough int) {
	if mock.UpdateResource3Func == nil {
		panic("ServerInterfaceMock.UpdateResource3Func: method is nil but ServerInterface.UpdateResource3 was just called")
	}
	callInfo := struct {
		W            http.ResponseWriter
		R            *http.Request
		PFallthrough int
	}{
		W:            w,
		R:            r,
		PFallthrough: pFallthrough,
	}
	mock.lockUpdateResource3.Lock()
	mock.calls.UpdateResource3 = append(mock.calls.UpdateResource3, callInfo)
	mock.lockUpdateResource3.Unlock()
	mock.UpdateResource3Func(w, r, pFallthrough)
}

// UpdateResource3Calls gets all the calls that were made to UpdateResource3.
// Check the length with:
//     len(mockedServerInterface.UpdateResource3Calls())
func (mock *ServerInterfaceMock) UpdateResource3Calls() []struct {
	W            http.ResponseWriter
	R            *http.Request
	PFallthrough int
} {
	var calls []struct {
		W            http.ResponseWriter
		R            *http.Request
		PFallthrough int
	}
	mock.lockUpdateResource3.RLock()
	calls = mock.calls.UpdateResource3
	mock.lockUpdateResource3.RUnlock()
	return calls
}
