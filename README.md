# Go Learning Practice

A comprehensive collection of Go programming concepts and examples organized by topic.

## Root Level (`main.go`)
- **If conditions** - Basic conditional statements
- **Statement; statement idiom** - Multiple statements on one line
- **Comma ok idiom** - Pattern for checking operation success
- **Switch case** - Switch statements and type switches
- **For loop** - All variations of Go's single loop construct

## Project Structure

### 📊 **Data Structures**
- **`array_slice/`** - Arrays vs slices, slice operations, capacity vs length
- **`maps/`** - Hash maps, key-value pairs, map operations
- **`structs/`** - Custom types, embedded structs, field promotion

### 🔧 **Functions & Methods**
- **`functions/`** - Function declarations, anonymous functions, callbacks
- **`generics/`** - Type parameters, type constraints, type sets with `~` operator
- **`interfaces_and_polymorphism/`** - Interface definitions, polymorphic behavior
- **`method-sets-1/`** - Value vs pointer receivers, method sets
- **`recursion/`** - Recursive function patterns

### 🚀 **Concurrency**
- **`concurrency/`** - Goroutines basics, concurrent execution
- **`channels/`** - Channel communication, directional channels
  - **`Directional Channels/`** - Send-only and receive-only channels
  - **`Select/`** - Select statement for channel operations
- **`mutex/`** - Race conditions, mutual exclusion, sync.Mutex
- **`workerPools/`** - Worker pool pattern, job distribution
- **`fan_in/`** - Fan-in pattern, merging multiple channels
- **`fan_out/`** - Fan-out pattern, distributing work
- **`context/`** - Context package for cancellation, timeouts, request-scoped values

### 🗄️ **Data & I/O**
- **`database_practice/`** - Database operations and queries
- **`working_with_json/`** - JSON marshaling/unmarshaling
- **`pointers/`** - Pointer semantics, memory addresses

### 🛠️ **Advanced Topics**
- **`garbage-collection-stack-and-heap/`** - Memory management, escape analysis
- **`bcrypt/`** - Password hashing and security

## Key Go Concepts Covered

### Memory Management
- **Stack vs Heap** - Value semantics (stack) vs pointer semantics (heap)
- **Escape Analysis** - Compiler optimization for memory allocation
- **Garbage Collection** - Automatic memory cleanup

### Concurrency Patterns
- **Goroutines** - Lightweight threads
- **Channels** - Communication between goroutines
- **Select** - Non-blocking channel operations
- **Worker Pools** - Efficient job processing
- **Fan-in/Fan-out** - Advanced channel patterns
- **Context** - Cancellation and timeout management

### Type System
- **Interfaces** - Contracts and polymorphism
- **Generics** - Type-safe code reuse
- **Method Sets** - Value vs pointer receivers
- **Struct Embedding** - Composition over inheritance

### Best Practices
- Prefer composition over inheritance
- Use interfaces for flexible design
- Handle errors explicitly
- Use context for cancellation
- Favor value semantics when possible
