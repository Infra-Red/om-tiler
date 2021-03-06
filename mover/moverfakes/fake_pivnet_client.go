// Code generated by counterfeiter. DO NOT EDIT.
package moverfakes

import (
	"context"
	"os"
	"sync"

	"github.com/starkandwayne/om-tiler/mover"
	"github.com/starkandwayne/om-tiler/pattern"
)

type FakePivnetClient struct {
	DownloadFileStub        func(context.Context, pattern.PivnetFile, string) (*os.File, error)
	downloadFileMutex       sync.RWMutex
	downloadFileArgsForCall []struct {
		arg1 context.Context
		arg2 pattern.PivnetFile
		arg3 string
	}
	downloadFileReturns struct {
		result1 *os.File
		result2 error
	}
	downloadFileReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) DownloadFile(arg1 context.Context, arg2 pattern.PivnetFile, arg3 string) (*os.File, error) {
	fake.downloadFileMutex.Lock()
	ret, specificReturn := fake.downloadFileReturnsOnCall[len(fake.downloadFileArgsForCall)]
	fake.downloadFileArgsForCall = append(fake.downloadFileArgsForCall, struct {
		arg1 context.Context
		arg2 pattern.PivnetFile
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("DownloadFile", []interface{}{arg1, arg2, arg3})
	fake.downloadFileMutex.Unlock()
	if fake.DownloadFileStub != nil {
		return fake.DownloadFileStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.downloadFileReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) DownloadFileCallCount() int {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return len(fake.downloadFileArgsForCall)
}

func (fake *FakePivnetClient) DownloadFileCalls(stub func(context.Context, pattern.PivnetFile, string) (*os.File, error)) {
	fake.downloadFileMutex.Lock()
	defer fake.downloadFileMutex.Unlock()
	fake.DownloadFileStub = stub
}

func (fake *FakePivnetClient) DownloadFileArgsForCall(i int) (context.Context, pattern.PivnetFile, string) {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	argsForCall := fake.downloadFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePivnetClient) DownloadFileReturns(result1 *os.File, result2 error) {
	fake.downloadFileMutex.Lock()
	defer fake.downloadFileMutex.Unlock()
	fake.DownloadFileStub = nil
	fake.downloadFileReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DownloadFileReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.downloadFileMutex.Lock()
	defer fake.downloadFileMutex.Unlock()
	fake.DownloadFileStub = nil
	if fake.downloadFileReturnsOnCall == nil {
		fake.downloadFileReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.downloadFileReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ mover.PivnetClient = new(FakePivnetClient)
