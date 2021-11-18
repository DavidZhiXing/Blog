#ifndef SQLITE_BTREE_H
#define SQLITE_BTREE_H

#define SQLITE_N_BTREE_META 16

#define BTREE_AUTOVACUUM_NONE 0
#define BTREE_AUTOVACUUM_FULL 1
#define BTREE_AUTOVACUUM_INCR 2

typedef struct Btree Btree;
typedef struct BtCursor BtCursor;
typedef struct BtreePayload BtreePayload;
typedef struct BtreeShared BtreeShared;

int sqlite3BtreeOpen(
  sqlite3_vfs *pVfs,       /* VFS to use with this b-tree */
  const char *zFilename,   /* Name of database file to open */
  sqlite3 *pDb,            /* Associated database connection */
  Btree **ppBtree,         /* Return open Btree* here */
  int flags,               /* Flags */
  int vfsFlags             /* Flags passed through to VFS open */
);

#define BTREE_OMIT_JOURNAL 0x0001  /* Do not create or use a rollback journal */
#define BTREE_NO_READLOCK 0x0002  /* Omit readlocks on readonly connections */
#define BTREE_MEMORY 0x0004       /* This is an in-memory DB */
#define BTREE_SINGLE 0x0008       /* The file contains at most 1 b-tree */

int sqlite3BtreeClose(Btree*);
int sqlite3BtreeSetCacheSize(Btree*,int);
int sqlite3BtreeSetSpillSize(Btree*,int);

int sqlite3BtreeSetPagerFlags(Btree*,unsigned);
int sqlite3BtreeSetPageSize(Btree*,int,int);
int sqlite3BtreeGetPageSize(Btree*);
Pgno sqlite3BtreeLastPage(Btree*);
Pgno sqlite3BtreeMaxPageCount(Btree*,int);
int sqlite3BtreeSecureDelete(Btree*,int);
int sqlite3BtreeGetReserve(Btree*);
#endif