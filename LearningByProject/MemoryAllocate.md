**Text Section** : The part that contains the binary instructions to be executed by the processor.

**Data Section** : Contains non-zero initialized static data.

**BSS** : Contains zero-initialized static data. Static data uninitialized in program is initialized 0 and goes here.

**Stack** : Contains the stack.

**Heap** : Contains the heap.

``` cpp
void* malloc(size_t size)
{
    void* block;  
    block = sbrk(size);
    if (block == (void*)-1)
        return NULL;
    return block;
}

union header{
    struct {
        size_t size;
        unsigned is_free;
        struct header_t* next;
    } s;
    ALIGN stub;
};
typedef union header header_t;

void* malloc(size_t size)
{
    size_t total_size;
    void* block;
    header_t* header;
    if (size == 0)
        return NULL;
    pthread_mutex_lock(&mutex);
    header = get_free_block(size);
    if(header){
        header->s.is_free = 0;
        pthread_mutex_unlock(&mutex);
        return (void*)(header + 1);
    }
    total_size = size + sizeof(header_t);
    block = sbrk(total_size);
    if (block == (void*)-1)
    {
        pthread_mutex_unlock(&mutex);
        return NULL;
    }
    header = (header_t*)block;
    header->s.size = size;
    header->s.is_free = 0;
    header->s.next = NULL;
    if(!head){
        head = header;
    }
    if(tail){
        tail->s.next = header;
    }
    tail = header;
    pthread_mutex_unlock(&mutex);
    return (void*)(header + 1);
}

void* get_free_block(size_t size)
{
    header_t* header;
    void *program_break = sbrk(0);
    header = head;
    while(header){
        if(header->s.is_free && header->s.size >= size){
            return header;
        }
        header = header->s.next;
    }
    return NULL;
}

void free(void* ptr)
{
    header_t* header;
    if(!ptr)
        return;
    pthread_mutex_lock(&mutex);
    header = (header_t*)ptr - 1;
    if((char*)ptr + header->s.size == program_break){
        if (head == tail){
            head = NULL;
            tail = NULL;
        }
        else{
            header_t* prev = head;
            while(prev->s.next != header){
                prev = prev->s.next;
            }
            prev->s.next = NULL;
            tail = prev;
        }
        sbrk(-header->s.size - sizeof(header_t));
        pthread_mutex_unlock(&mutex);
        return;
    }
    header->s.is_free = 1;
    pthread_mutex_unlock(&mutex);
}
```
calloc
``` cpp
void* calloc(size_t nmemb, size_t size)
{
    void* ptr;
    size_t total_size;
    total_size = nmemb * size;
    ptr = malloc(total_size);
    if(!ptr)
        return NULL;
    memset(ptr, 0, total_size);
    return ptr;
}
```
realloc
``` cpp
void realloc(void *block, size_t size) {
    header_t* header;
    void *ret;
    if(!block || !size)
        return malloc(size);
    header = (header_t*)block - 1;
    if(header->s.size >= size){
        return block;
    }
    ret = malloc(size);
    if(ret){
        memcpy(ret, block, header->s.size);
        free(block);
    }
    return ret;
}
```

### Compliling and using our memory allocator
``` cpp
$ gcc -o memalloc.so -shared -fPIC memalloc.c
$ export LD_PRELOAD=./memalloc.so
$ ls
``` cpp
