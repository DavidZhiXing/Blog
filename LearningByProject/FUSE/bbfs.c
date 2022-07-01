// All the paths are relative to the root of the mounted filesystem.
// In order to get to the underlying filesystem, I need to have the mountpoint.
// I will save it away early oon in main(), and then whenever I need a path for something
// I will call this to constuct it.

static void bb_fullpath(char fpath[PATH_MAX], const char *path)
{
    strcpy(fpath, BB_DATA->rootdir);
    strncat(fpath, path, PATH_MAX); // ridiculously long paths will cause problems
    log_msg("    fullpath: \"%s\"\n", fpath);
}

// prototypes for all these functions, and the C-style comments, come from /usr/include/fuse.h
// get file attributes
// similar to stat(). The 'st_dev' and 'st_blksize' fields are ignored. The 'st_ino' field is ignored except for on the MacOSX platform where it is used to disambiguate between files with the same name.

int bb_getattr(const char *path, struct stat *statbuf)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_getattr(path=\"%s\", statbuf=0x%08x)\n", path, statbuf);
    bb_fullpath(fpath, path);
    retstat = lstat(fpath, statbuf);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// read the target of a symbolic link
// the buffer should be filled with a null terminated string.
// the buffer size argument includes the space for the terminating null character.
// if the linkname is too long to fit in the buffer, it should be truncated. The return
// value should be 0 for success.
// note:
// the system readlink() will turcated and lose the terminating null. so, the size passed to the system readlink() must be one less than the size of the buffer.
// must be  one  bb_readlink() code by bernardo F Costa's code.
int bb_readlink(const char *path, char *link, size_t size)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_readlink(path=\"%s\", link=\"%s\", size=%d)\n", path, link, size);
    bb_fullpath(fpath, path);
    retstat = readlink(fpath, link, size - 1);
    if (retstat < 0)
        retstat = -errno;
    else
    {
        link[retstat] = '\0';
        retstat = 0;
    }
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// Create a file node
// There is no create() operation. mknod() will be called for creation of 
// all non-directory, non-symlink nodes.
// shouldn't that comment be "if" there is no... ?
int bb_mknod(const char *path, mode_t mode, dev_t dev)
{
    int restat = 0;
    char fpath[PATH_MAX];

    log_msg("bb_mknod(path=\"%s\", mode=0%3o, dev=%lld)\n", path, mode, dev);
    bb_fullpath(fpath, path);

    // on linux this could just be 'mknod(path, mode, dev)' but this tries to be more portable
    // by hornoring the quote in the linux mknod man page stating the only portable use of mknod()
    // is to make a fifo, but saving never actually be used for that.
    if (S_ISREG(mode)) {
        restat = log_syscall ("open", open(fpath, 0_CREAT | 0_EXCL | 0_WRONLY, mode), 0);
        if (restat >= 0) {
            restat = log_syscall ("close", close(restat), 0);
        }
    } else if (S_ISFIFO(mode)) {
        restat = log_syscall ("mkfifo", mkfifo(fpath, mode), 0);
    } else {
        restat = log_syscall ("mknod", mknod(fpath, mode, dev), 0);
    }
    return restat;
}

// create a directory
int bb_mkdir(const char *path, mode_t mode)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_mkdir(path=\"%s\", mode=0%3o)\n", path, mode);
    bb_fullpath(fpath, path);
    retstat = log_syscall("mkdir", mkdir(fpath, mode), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// remove a file
int bb_unlink(const char *path)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_unlink(path=\"%s\")\n", path);
    bb_fullpath(fpath, path);
    retstat = log_syscall("unlink", unlink(fpath), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// remove a directory
int bb_rmdir(const char *path)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_rmdir(path=\"%s\")\n", path);
    bb_fullpath(fpath, path);
    retstat = log_syscall("rmdir", rmdir(fpath), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// create a symbolic link
// the parameters here are a little bit confusing, but dow correspond
// to the symlink() system call. The 'path' is where the link points, while
// the 'line' is the link itself. So we need to leave the path unatered, but insert the link into
// the mounted directory.
int bb_symlink(const char *path, const char *link)
{
    int retstat = 0;
    char flink[PATH_MAX];
    log_msg("bb_symlink(path=\"%s\", link=\"%s\")\n", path, link);
    bb_fullpath(flink, link);
    retstat = log_syscall("symlink", symlink(path, flink), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// rename a file
int bb_rename(const char *path, const char *newpath)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    char fnewpath[PATH_MAX];
    log_msg("bb_rename(path=\"%s\", newpath=\"%s\")\n", path, newpath);
    bb_fullpath(fpath, path);
    bb_fullpath(fnewpath, newpath);
    retstat = log_syscall("rename", rename(fpath, fnewpath), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// create a hard link to a file
int bb_link(const char *path, const char *newpath)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    char fnewpath[PATH_MAX];
    log_msg("bb_link(path=\"%s\", newpath=\"%s\")\n", path, newpath);
    bb_fullpath(fpath, path);
    bb_fullpath(fnewpath, newpath);
    retstat = log_syscall("link", link(fpath, fnewpath), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// change the permission bits of a file
int bb_chmod(const char *path, mode_t mode)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_chmod(path=\"%s\", mode=0%03o)\n", path, mode);
    bb_fullpath(fpath, path);
    retstat = log_syscall("chmod", chmod(fpath, mode), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// change the owner and group of a file
int bb_chown(const char *path, uid_t uid, gid_t gid)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_chown(path=\"%s\", uid=%d, gid=%d)\n", path, uid, gid);
    bb_fullpath(fpath, path);
    retstat = log_syscall("chown", chown(fpath, uid, gid), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// change the size of a file
int bb_truncate(const char* path, off_t size)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_truncate(path=\"%s\", size=%lld)\n", path, size);
    bb_fullpath(fpath, path);
    retstat = log_syscall("truncate", truncate(fpath, size), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// change the access and/or  modification times of a file
// note -- I will want to  change this as soon as 2.6 is in debian testing
int bb_utime(const char *path, struct utimbuf *ubuf)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_utime(path=\"%s\", ubuf=0x%08x)\n", path, ubuf);
    bb_fullpath(fpath, path);
    retstat = log_syscall("utime", utime(fpath, ubuf), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// file open  operation
// no creation, or truncation flags (O_CREAT, O_EXCL, O_TRUNC)
// will be used. Open should check if the operation is  permitted for the given
// flags. Optionally open may also return an arbitray filehandle in the fuse_file_info structure,
// which will be passed to all file operations.
int bb_open(const char *path, struct fuse_file_info *fi)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_open(path=\"%s\", fi=0x%08x)\n", path, fi);
    bb_fullpath(fpath, path);
    retstat = log_syscall("open", open(fpath, fi->flags), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// read data from an open file
// read should return exactly the number of bytes requested except on EOF or error,
// otherwise the data will be substitued with zeroes. An exception to this is when the 
// 'direct_io' mount option is specified, in which case the return value of the read system call
// will reflect the return value of this operation.
int bb_read(const char *path, char *buf, size_t size, off_t offset, struct fuse_file_info *fi)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_read(path=\"%s\", buf=0x%08x, size=%d, offset=%lld, fi=0x%08x)\n", path, buf, size, offset, fi);
    bb_fullpath(fpath, path);
    retstat = log_syscall("read", pread(fi->fh, buf, size, offset), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// write data to an open file

// Write should retrun exactly the number of bytes requested except on error.
// An exception to this is when the 'direct_io' mount option is specified
int bb_write(const char *path, const char *buf, size_t size, off_t offset, struct fuse_file_info *fi)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_write(path=\"%s\", buf=0x%08x, size=%d, offset=%lld, fi=0x%08x)\n", path, buf, size, offset, fi);
    bb_fullpath(fpath, path);
    retstat = log_syscall("write", pwrite(fi->fh, buf, size, offset), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// Get file system statistics
int bb_stats(const char *path, struct statvfs *statv){
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_stats(path=\"%s\", statv=0x%08x)\n", path, statv);
    bb_fullpath(fpath, path);
    retstat = log_syscall("statvfs", statvfs(fpath, statv), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// possibly flush cached data
// BIG NOTE: This is not equivalent to fsync(). It's not a request to sync dirty
// data. It's actually a request to flush data to disk. A flush is a request to
// writeback data to disk.
int bb_flush(const char *path, struct fuse_file_info *fi)
{
    int retstat = 0;
    log_msg("bb_flush(path=\"%s\", fi=0x%08x)\n", path, fi);
    // no need to get fpath on this one, since I work from fi->fh not the path
    retstat = log_syscall("flush", fsync(fi->fh), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// release an open file
// release is called when there are no more references to an open file: all file
// descriptions are closed and all memory mappings are unmapped.
// for every open() call there will be exactly one release() call with the same flags and file
// descriptor. It is possible to have a file opened more than once, in which case only the last
// release will mean, that no more reads/writes will happen on the file. The return value of release is 
// ignored.
int bb_release(const char *path, struct fuse_file_info *fi)
{
    int retstat = 0;
    log_msg("bb_release(path=\"%s\", fi=0x%08x)\n", path, fi);
    retstat = log_syscall("release", close(fi->fh), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// sychronize file contents
// if the datasync parameter is non-zero, then only the user data
// should be fulushed, not the meta data.
int bb_fsync(const char *path, int datasync, struct fuse_file_info *fi)
{
    int retstat = 0;
    log_msg("bb_fsync(path=\"%s\", datasync=%d, fi=0x%08x)\n", path, datasync, fi);
    retstat = log_syscall("fsync", fsync(fi->fh), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// set extended attributes
int bb_setxattr(const char *path, const char *name, const char *value, size_t size, int flags)
{
    char fpath[PATH_MAX];

    log_msg("bb_setxattr(path=\"%s\", name=\"%s\", value=\"%s\", size=%d, flags=0x%08x)\n", path, name, value, size, flags);
    bb_fullpath(fpath, path);
    int retstat = log_syscall("setxattr", lsetxattr(fpath, name, value, size, flags), 0);
    return retstat;
}

int bb_getxattr(const char *path, const char *name, char *value, size_t size)
{
    char fpath[PATH_MAX];
    int retstat = 0;
    log_msg("bb_getxattr(path=\"%s\", name=\"%s\", value=\"%s\", size=%d)\n", path, name, value, size);
    bb_fullpath(fpath, path);
    retstat = log_syscall("getxattr", lgetxattr(fpath, name, value, size), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

int bb_listxattr(const char *path, char *list, size_t size)
{
    char fpath[PATH_MAX];
    int retstat = 0;
    log_msg("bb_listxattr(path=\"%s\", list=0x%08x, size=%d)\n", path, list, size);
    bb_fullpath(fpath, path);
    retstat = log_syscall("listxattr", llistxattr(fpath, list, size), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

int bb_removexattr(const char *path, const char *name) 
{
    char fpath[PATH_MAX];
    int retstat = 0;
    log_msg("bb_removexattr(path=\"%s\", name=\"%s\")\n", path, name);
    bb_fullpath(fpath, path);
    retstat = log_syscall("removexattr", lremovexattr(fpath, name), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

int bb_opendir(const char *path, struct fuse_file_info *fi)
{
    int retstat = 0;
    char fpath[PATH_MAX];
    log_msg("bb_opendir(path=\"%s\", fi=0x%08x)\n", path, fi);
    bb_fullpath(fpath, path);
    DIR *dp;
    dp = opendir(fpath);
    if (dp == NULL)
        retstat = -errno;
    fi->fh = (intptr_t) dp;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// read directory
// this supersedes the old getdir() interface. new applications should use this.

// the filesystem may choose between two modes of operation:
// 1) the directory implementation ignores the offset parameter, and 
//    passes zero to the filler function's offset。The filler
//    function will not return '1' (unless an error happens), so the
//    whole directory is read in a single readdir operation. This
//    works just like the old getdir() method.

// 2) the directory implementation keeps track of the offsets of the
//    directory entries. It uses the offset parameter and always
//    passes non-zero offset to the filler function. When the buffer
//    is full (or an error happens), the filler function will return
//    '1'.
int bb_readdir(const char *path, void *buf, fuse_fill_dir_t filler, off_t offset, struct fuse_file_info *fi)
{
    int retstat = 0;
    log_msg("bb_readdir(path=\"%s\", buf=0x%08x, filler=0x%08x, offset=%lld, fi=0x%08x)\n", path, buf, filler, offset, fi);
    DIR *dp;
    struct dirent *de;
    dp = (DIR *) (uintptr_t) fi->fh;
    de = readdir(dp);
    if (de == 0)
        retstat = -errno;
    while (de != 0) {
        if (filler(buf, de->d_name, NULL, 0) != 0)
            break;
        de = readdir(dp);
    }
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// release directory
int bb_releasedir(const char *path, struct fuse_file_info *fi)
{
    int retstat = 0;
    log_msg("bb_releasedir(path=\"%s\", fi=0x%08x)\n", path, fi);
    retstat = log_syscall("releasedir", closedir((DIR *) (uintptr_t) fi->fh), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// synchronize directory contents
int bb_fsyncdir(const char *path, int datasync, struct fuse_file_info *fi)
{
    int retstat = 0;
    log_msg("bb_fsyncdir(path=\"%s\", datasync=%d, fi=0x%08x)\n", path, datasync, fi);
    retstat = log_syscall("fsyncdir", fsync(((DIR *) (uintptr_t) fi->fh)->fd), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// initialize filesystem
int bb_init(struct fuse_conn_info *conn, struct fuse_config *cfg)
{
    log_msg("bb_init()\n");
    return 0;
}

// clean up filesystem
void bb_destroy(void *userdata)
{
    log_msg("bb_destroy(userdata=0x%08x)\n", userdata);
}

// check file access permissions
int bb_access(const char *path, int mask)
{
    int retstat = 0;
    char fpath[PATH_MAX];

    log_msg("bb_access(path=\"%s\", mask=0%o)\n", path, mask);
    bb_fullpath(fpath, path);

    retstat = log_syscall("access", access(fpath, mask), 0);
    if (retstat < 0)
        retstat = -errno;
    return retstat;
}

// create and open a file
int bb_ftruncate(const char *path, off_t offset, struct fuse_file_info *fi)
{
    int retstat = 0;
    log_msg("bb_ftruncate(path=\"%s\", offset=%lld, fi=0x%08x)\n", path, offset, fi);
    retstat = log_syscall("ftruncate", ftruncate(fi->fh, offset), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

// get attibutes from an open file
int bb_fgetattr(const char *path, struct stat *statbuf, struct fuse_file_info *fi)
{
    int retstat = 0;
    log_msg("bb_fgetattr(path=\"%s\", statbuf=0x%08x, fi=0x%08x)\n", path, statbuf, fi);
    retstat = log_syscall("fgetattr", fstat(fi->fh, statbuf), 0);
    if (retstat < 0)
        retstat = -errno;
    log_msg("    retstat = %d\n", retstat);
    return retstat;
}

struct fuse_operations bb_oper = {
    .getattr = bb_getattr,
    .readlink = bb_readlink,
    .getdir = NULL,
    .mknod = bb_mknod,
    .mkdir = bb_mkdir,
    .unlink = bb_unlink,
    .rmdir = bb_rmdir,
    .symlink = bb_symlink,
    .rename = bb_rename,
    .link = bb_link,
    .chmod = bb_chmod,
    .chown = bb_chown,
    .truncate = bb_truncate,
    .utime = bb_utime,
    .open = bb_open,
    .read = bb_read,
    .write = bb_write,
    .statfs = bb_statfs,
    .flush = bb_flush,
    .release = bb_release,
    .fsync = bb_fsync,
    .setxattr = bb_setxattr,
    .getxattr = bb_getxattr,
    .listxattr = bb_listxattr,
    .removexattr = bb_removexattr,
    .opendir = bb_opendir,
    .readdir = bb_readdir,
    .releasedir = bb_releasedir,
    .fsyncdir = bb_fsyncdir,
    .init = bb_init,
    .destroy = bb_destroy,
    .access = bb_access,
    .ftruncate = bb_ftruncate,
    .fgetattr = bb_fgetattr,
    .lock = NULL,
    .utimens = NULL,
    .bmap = NULL,
    .ioctl = NULL,
    .poll = NULL,
    .write_buf = NULL,
    .read_buf = NULL,
    .flock = NULL,
    .fallocate = NULL
};

void bb_usage()
{
    fprintf(stderr, "usage: bbfs [FUSE and mount options] rootDir mountPoint\n");
    abort();
}

int main(int argc, char *argv[])
{
    int fuse_stat;
    struct bb_state *bb_data;

    if (getuid() == 0 || (geteuid() == 0)) {
        fprintf(stderr, "bbfs should not be run as root.\n");
        exit(1);
    }

    // see which versions of fuse we're running
    fprintf(stderr, "Fuse library version %d.%d\n", FUSE_MAJOR_VERSION, FUSE_MINOR_VERSION);
    if (argc < 3 || argv[argc - 2][0] == '-') {
        bb_usage();
    }

    bb_data = malloc(sizeof(struct bb_state));
    if (bb_data == NULL) {
        fprintf(stderr, "main: unable to allocate memory\n");
        exit(1);
    }

    bb_data->rootdir = realpath(argv[argc - 1], NULL);
    argv[argc - 1] = argv[argc - 2];
    argv[argc - 2] = NULL;
    argc--;
    
    bb_data->logfile = log_open();

    fprintf(stderr, "main: bb_data->rootdir = %s\n", bb_data->rootdir);
    fuse_stat = fuse_main(argc, argv, &bb_oper, bb_data);
    fprintf(stderr, "main: fuse_main returned %d\n", fuse_stat);

    return fuse_stat;
}