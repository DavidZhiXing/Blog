### basic main

``` cpp
int main() {
    // load config files, if any.

    // run command loop.
    lsh_loop();

    // perform any shutdown/cleanup.

    return EXIT_SUCCESS;
}
```

### basic loop

``` cpp
void lsh_loop() {
    char *line;
    char **args;
    char status;

    do {
        printf("> ");
        line = lsh_read_line();
        args = lsh_split_line(line);
        status = lsh_execute(args);

        free(line);
        free(args);
    } while (status);
}
```

### reading a line

``` cpp
char *lsh_read_line(void) {
    int bufsize = LSH_RL_BUFSIZE;
    int position = 0;
    char *buffer = malloc(sizeof(char) * bufsize);
    int c;

    if (!buffer) {
        fprintf(stderr, "lsh: allocation error\n");
        exit(EXIT_FAILURE);
    }

    while (l) {
        // read a charactor
        c = getchar();

        // if we hit EOF, replace it with a null character and return.
        if (c == EOF || c == '\n') {
            buffer[position] = '\0';
            return buffer;
        } else {
            buffer[position] = c;
        }
        position++;

        // If we have exceeded the buffer, reallocate.
        if (position >= bufsize) {
            bufsize += LSH_RL_BUFSIZE;
            buffer = realloc(buffer, bufsize);
            if (!buffer) {
                fprintf(stderr, "lsh: allocation error\n");
                exit(EXIT_FAILURE);
            }
        }
    }
}

// using getline()
char *lsh_read_line(void) {
    char *line = NULL;
    size_t line_size;

    if (getline(&line, &line_size, stdin) == -1) {
        exit(EXIT_FAILURE);
    }

    return line;
}
```

### parsing the line

``` cpp
char **lsh_split_line(char *line) {
    int bufsize = LSH_TOKEN_BUFSIZE, position = 0;
    char **tokens = malloc(bufsize * sizeof(char*));
    char *token, **tokens_backup;

    if (!tokens) {
        fprintf(stderr, "lsh: allocation error\n");
        exit(EXIT_FAILURE);
    }

    token = strtok(line, LSH_TOKEN_DELIM);
    while (token != NULL) {
        tokens[position] = token;
        position++;
        if (position >= bufsize) {
            bufsize += LSH_TOKEN_BUFSIZE;
            tokens_backup = tokens;
            tokens = realloc(tokens, bufsize * sizeof(char*));
            if (!tokens) {
                free(tokens_backup);
                fprintf(stderr, "lsh: allocation error\n");
                exit(EXIT_FAILURE);
            }
        }
        token = strtok(NULL, LSH_TOKEN_DELIM);
    }
    tokens[position] = NULL;
    return tokens;
}
```

### start process

``` cpp
int lsh_launch(char **args) {
    pid_t pid, wpid;
    int status;

    pid = fork();
    if (pid == 0) {
        // Child process
        if (execvp(args[0], args) == -1) {
            perror("lsh");
            exit(EXIT_FAILURE);
        }
    } else if (pid < 0) {
        // Error forking
        perror("lsh");
    } else {
        // Parent process
        do {
            wpid = waitpid(pid, &status, WUNTRACED);
        } while (!WIFEXITED(status) && !WIFSIGNALED(status));
    }
}
```

### shell building

``` cpp
int lsh_cd(char **args) {
    if (args[1] == NULL) {
        fprintf(stderr, "lsh: expected argument to \"cd\"\n");
    } else {
        if (chdir(args[1]) != 0) {
            perror("lsh");
        }
    }
    return 1;
}

int lsh_help(char **args) {
    int i;
    for (i = 0; i < lsh_num_builtins(); i++) {
        printf("%s\n", lsh_builtins[i]);
    }
    return 1;
}

int lsh_exit(char **args) {
    return 0;
}

char *buildin_str[] = {
    "cd",
    "help",
    "exit"
};

int (*buildin_func[]) (char **) = {
    &lsh_cd,
    &lsh_help,
    &lsh_exit
};


```

### putting together buitins and processes

``` cpp
int lsh_execute(char **args) {
    int i;
    if (args[0] == NULL) {
        return 1;
    }
    for (i = 0; i < lsh_num_builtins(); i++) {
        if (strcmp(args[0], buildin_str[i]) == 0) {
            return (*buildin_func[i])(args);
        }
    }
    return lsh_launch(args);
}
```

### putting it all together

- #include <sys/wait.h>
    - waitpid() and associated macros
- #include <unistd.h>
    - chdir()
    - fork()
    - exec()
    - pid_t
- #include <stdlib.h>
    - malloc()
    - free()
    - exit()
    - realloc()
    - execvp()
    - EXIT_FAILURE, EXIT_SUCCESS
- #include <stdio.h>
    - fprintf()
    - perror()
    - printf()
    - getchar()
    - stderr
- #include <string.h>
    - strtok()
    - strcmp()
    
### wrapup