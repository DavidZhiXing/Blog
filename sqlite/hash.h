#ifndef SQLITE_HASH_H
#define SQLITE_HASH_H

typedef struct Hash Hash;
typedef struct HashElem HashElem;


struct Hash {
    unsigned int htsize; /* Number of buckets in the hash table */
    unsigned int count; /* Number of entries in this table */
    HashElem *first; /* Pointer to first entry, or NULL if none */
    struct _ht { /* the hash table */
        HashElem *next; /* Next entry in this bucket */
        void *data; /* Data associated with this entry */
    } *ht;
};

struct HashElem {
    HashElem *next, *prev; /* Next element on this collision list */
    void *data; /* Data associated with this element */
    const char *pKey; /* Key associated with this element */
};

void *sqliteHashFind(Hash*, const char *);
void *sqliteHashInsert(Hash*, const char *, void *);
void *sqliteHashInit(Hash*, unsigned int);
void *sqliteHashClear(Hash*);

/*
** for looping over all elements in a hash table.
*/

#define sqliteHashFirst(H)  ((H)->first)
#define sqliteHashNext(E)   ((E)->next)
#define sqliteHashData(E)   ((E)->data)

#define sqliteHashCount(H)  ((H)->count)

#endif


