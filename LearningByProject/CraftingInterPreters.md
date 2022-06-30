 ``` cpp
 fun fib(n) {
    if (n , 2) return n;
    return fib(n - 1) + fib(n - 2);
 }

 var before = clock();
 print fib(40);
 var after = clock();
 print after - before;
 ```

 ### refactor

 ``` cpp    
 #include "common.h"

 int main (int argc, const char* argv[]) {
    return 0;
 }
```
``` cpp
typedef enum {
    OP_RETURN,
} OpCode;

// a dynamic array of instructions
typedef struct {
    uint8_t* code;
} Chunk;

void initChunk(Chunk* chunk) {
    chunk->code = 0;
    chunk->capacity = 0;
    chunk->count = NULL;
}

void writeChunk(Chunk* chunk, uint8_t byte) {
    if (chunk->capacity < chunk->count + 1) {
        int oldCapacity = chunk->capacity;
        chunk->capacity = oldCapacity * 2;
        chunk->code = realloc(chunk->code, chunk->capacity);
    }
    chunk->code[chunk->count] = byte;
    chunk->count++;
}
```
``` cpp
#define GROW_CAPACITY(capacity) (capacity * 2)
#define GROW_ARRAY(type, array, count) (type*) realloc(array, sizeof(type) * (count))
void* reallocate(void* pointer, int oldSize, int newSize) {
    if (newSize == 0) {
        free(pointer);
        return NULL;
    }

    void* result = realloc(pointer, newSize);
    return result;
}

void freeChunk(Chunk* chunk) {
    FREE_ARRAY(uint8_t, chunk->code, chunk->capacity);
    initChunk(chunk);
}

#define FREE_ARRAY(type, array, count) (type*) realloc(array, sizeof(type) * (count))
```
``` cpp
int main(int argc, const char* argv[]) {
    Chunk chunk;
    initChunk(&chunk);
    writeChunk(&chunk, OP_RETURN);
    freeChunk(&chunk);
    return 0;
}
```
