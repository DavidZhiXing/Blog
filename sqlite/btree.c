/*
** This file implements a external (disk-based) database using BTree.
** For a detailed description of BTree, refer to:
**
**     Donald E. Knuth, THE ART OF COMPUTER PROGRAMMING, Volume 3:
**     "Sorting and Searching", pages 473-480. Addison-Wesley
**     Publishing Company, Reading, Massachusetts.
**
** The basic idea is that each page of the file contains N database
** entries and N+1 pointers to subpages.  (i.e. Each page has one
** more pointer than the number of entries it contains.  The first
** pointer of the first page is unused.  All subsequent pages can be
** determined by counting backward from the first pointer.  The
** pointer to the i-th page is the address of the (i-1)st entry in
** the (i-1)st page.  The first page is page 1, the second is page
** 2, and so forth.
**
** -----------------------------------------------------------------------------   
**  |  Ptr(0) | Key(0) | Ptr(1) | Key(1) | Ptr(2) | Key(2) | ... |Key(N)|Ptr(N+1)|
** -----------------------------------------------------------------------------
**
** All of the keys on the page that Ptr(0) points to have values less
** than Key(0).  All of the keys on page Ptr(1) and its subpages have
** values greater than Key(0) and less than Key(1).  All of the keys
** on Ptr(N) and its subpages have values greater than Key(N-1) and
** less than Key(N).  And so forth.
**
** Finding a particular key requires reading O(log(M)) pages from the
** disk where M is the number of entries in the tree.
**
** In this implementation, a single file can hold one or more separate
** BTrees.  Each BTree is identified by the index of its root page.
** The key and data for any entry are combined to form the "payload".
** Up to MX_PAGE_SIZE bytes of the payload can be carried directly
** on the database page.  If the payload is larger than this, then  
** the payload is stored on an overflow page.  The payload for an
** enttry and the preceding pointer are combined to form a "Cell".
** Each page has a small header which contains the Ptr(N+1) pointer.
**
** The first page of the file contains a magic string used to verify
** that the file is a valid BTree database, a pointer to a list of unused
** pages, in the file, and some meta information. The root of the first
** BTree begins on page 2 of the file. (Pages are numbered beginning with
** 1, not 0). Thus a minimum database contains 2 pages.
*/


typedef struct Btree Btree;
typedef struct BtCursor BtCursor;
typedef struct BtShared BtShared;

/*
** forward declarations of structure used by this file
*/
typedef struct PageOnee PageOnee;
typedef struct MemPage MemPage;
typedef struct Pgno Pgno;
typedef struct PageHdr PageHdr;
typedef struct Cell Cell;
typedef struct CellHeader CellHeader;
typedef struct FreeBlk FreeBlk;
typedef struct OverFlowPage OverFlowPage;
typedef struct FreelistInfo FreelistInfo;

int sqliteBtreeOpen(
  const char *zFilename,   /* Name of database file to open */
  Btree **,                /* Return open Btree* here */
  int flags,               /* Flags */
  int vfsFlags              /* Flags passed through to VFS open */
);

int sqliteBtreeClose(Btree*);
int sqliteBtreeSetCacheSize(Btree*, int);


int sqliteBtreeCursor(Btree*, int, int, BtCursor**);
int sqliteBtreeMoveto(BtCursor*, int, const void *, int, int*);
int sqliteBtreeDelete(BtCursor*);
int sqliteBtreeInsert(BtCursor*, const void *, int, int, int*);
int sqliteBtreeFirst(BtCursor*, int*);
int sqliteBtreeLast(BtCursor*, int*);
int sqliteBtreeNext(BtCursor*, int*);
int sqliteBtreeKey(BtCursor*, const void **, int*);
int sqliteBtreeKeySize(BtCursor*, int*);
int sqliteBtreeKeyCompare(BtCursor*, const void *, int, int*);

int sqliteBtreeCreateTable(Btree*, int*, int, int, int);
int sqliteBtreeDropTable(Btree*, int, int);
int sqliteBtreeClearTable(Btree*, int);
int sqliteBtreeCreateIndex(Btree*, int, int, int, int);

int sqliteBtreeGetMeta(Btree*, int, int*);
int sqliteBtreeUpdateMeta(Btree*, int*);


int sqliteBtreePageDump(Btree*, int, int);
int sqliteBtreeCursorDump(BtCursor*, int);
struct Pager* sqliteBtreePager(Btree*);

int sqliteBtreeData(BtCursor*, int, char**);
int sqliteBteeDataSize(BtCursor*);
int sqliteBtreeCloseCursor(BtCursor*);

int sqliteBtreeBeginTrans(Btree*, int);
int sqliteBtreeCommit(Btree*);
int sqliteBtreeRollback(Btree*);
int sqliteBtreeBeginCkpt(Btree*);
int sqliteBtreeEndCkpt(Btree*);
int sqliteBtreeRollbackCkpt(Btree*);


struct PageOne{
    char zMagic[8];
    int iVersion;
    Pgno freelist;
    int nFree;
    int aMeta[SQLITE_N_BTREE_META - 1];
};



