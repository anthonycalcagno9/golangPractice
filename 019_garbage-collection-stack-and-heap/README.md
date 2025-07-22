# Go Memory Management: Stack vs Heap

## The Stack
- Typically, when functions use value semantics, the data is stored on the stack
- This is good, the stack is fast and efficient and **DOES NOT INVOLVE GARBAGE COLLECTION**

## The Heap
- When functions use pointer semantics, the data must be stored on the heap
- That data will need to be used again so it **ESCAPES THE STACK AND GOES TO THE HEAP**
- The heap is slower and requires garbage collection to clean up unused data

## Escape Analysis
- Go Compiler uses escape analysis to determine whether a variable should be allocated on the stack or the heap
- Can use flag `-gcflags -m` alongside `go build` or `go run` to see escape analysis results
- Example: `go run -gcflags -m main.go`

## Inlining Decisions
- The Go compiler can inline functions, which means it replaces the function call with the function body
- This can improve performance by avoiding the overhead of a function call
- Inlining is done automatically by the compiler based on heuristics
- You can use the `-gcflags -m` flag to see inlining decisions as well

## Key Concepts Summary

| Aspect | Stack | Heap |
|--------|-------|------|
| **Speed** | Fast | Slower |
| **Garbage Collection** | No | Yes |
| **Used for** | Value semantics | Pointer semantics |
| **Memory Management** | Automatic (function scope) | Requires GC cleanup |

## Performance Tips
- Prefer value semantics when possible to keep data on the stack
- Use pointer semantics only when necessary (sharing data, large structs)
- The Go compiler is smart about escape analysis - trust it but understand it
