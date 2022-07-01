FILE *log_open()
{
    FILE *log_file = fopen("log.txt", "w");
    if (log_file == NULL)
    {
        printf("Error opening log file\n");
        exit(1);
    }

    setvbuf(log_file, NULL, _IONBF, 0);

    return log_file;
}

void log_msg(const char * format, ...)
{
    va_list args;
    va_start(args, format);
    vfprintf(log_file, format, args);
    va_end(args);
}
{
    fprintf(log_file, "%s\n", msg);
}

int log_err(char *func)
{
    int ret = -errno;

    log_msg("%s: %s\n", func, strerror(errno));

    return ret;
}

void log_fuse_context(struct fuse_context *ctx)
{
    log_msg("fuse_context: uid=%d gid=%d pid=%d\n", ctx->uid, ctx->gid, ctx->pid);
}

void log_conn(struct fuse_conn_info *conn)
{
    log_msg("conn: max_read=%d max_write=%d max_readahead=%d max_writeback=%d\n",
            conn->max_read, conn->max_write, conn->max_readahead, conn->max_writeback);
}

void log_fi(struct fuse_file_info *fi)
{
    log_msg("fi: flags=%d fh=%d writepage=%d\n", fi->flags, fi->fh, fi->writepage);
}

int log_syscall(char *func, int retstat, int min_ret)
{
    if (retstat < min_ret)
    {
        log_msg("%s: %s\n", func, strerror(retstat));
        return -retstat;
    }
    else
    {
        log_msg("%s: %d\n", func, retstat);
        return retstat;
    }
}

void log_retstat(char *func, int retstat)
{
    log_msg("%s: %d\n", func, retstat);
}

void log_stat(struct stat *si) {
    log_msg("si: st_dev=%d st_ino=%d st_mode=%d st_nlink=%d st_uid=%d st_gid=%d st_rdev=%d st_size=%d st_blksize=%d st_blocks=%d st_atime=%d st_mtime=%d st_ctime=%d\n",
            si->st_dev, si->st_ino, si->st_mode, si->st_nlink, si->st_uid, si->st_gid, si->st_rdev, si->st_size, si->st_blksize, si->st_blocks, si->st_atime, si->st_mtime, si->st_ctime);
}

void log_utime(struct utimbuf *ut) {
    log_msg("ut: actime=%d modtime=%d\n", ut->actime, ut->modtime);
}