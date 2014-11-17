#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define	G16L(p)	((p)[0]|((p)[1]<<8))
#define	G16B(p)	((p)[0]<<8|((p)[1]))
#define	G32L(p) 	((p)[0]|((p)[1]<<8)|((p)[2]<<16)|((p)[3]<<24))
#define	G32B(p)	(((p)[0]<<24)|((p)[1]<<16)|((p)[2]<<8)|(p)[3])
#define	P16L(p,v)	(p)[0]=(v);(p)[1]=(v)>>8
#define	P16B(p,v)	(p)[0]=(v)>>8;(p)[1]=(v)
#define	P32L(p,v)	(p)[0]=(v);(p)[1]=(v)>>8;(p)[2]=(v)>>16;(p)[3]=(v)>>24
#define	P32B(p,v)	(p)[0]=(v)>>24;(p)[1]=(v)>>16;(p)[2]=(v)>>8;(p)[3]=(v)

#define NPASSES (10 * 1000 * 1000)

void
sp(char *msg, struct timeval *s);

int 
main() {
	bench(NPASSES);
	return 0;
}

int 
bench(int n) {
	struct timeval s;

	int i;
	int v = 0x41424344;

	char *b = malloc(n*4);

	gettimeofday(&s, NULL);
	for (i = 0; i < n; i++) 
		P32B(b+i*4, v);
	sp("P32B",&s);

	gettimeofday(&s, NULL);
	for (i = 0; i < n; i++) 
		P32L(b+i*4, v);
	sp("P32L",&s);

	gettimeofday(&s, NULL);
	for (i = 0; i < n; i++) 
		G32B(b+i*4);
	sp("G32B",&s);

	gettimeofday(&s, NULL);
	for (i = 0; i < n; i++) 
		G32L(b+i*4);
	sp("G32L",&s);

}

void
sp(char *msg, struct timeval *s) {
	struct timeval e;
	gettimeofday(&e, NULL);

	long long delta;
	delta = (e.tv_usec - s->tv_usec) / 1000;
	delta += (e.tv_sec - s->tv_sec) * 1000;

	printf("c: %s: %d ops took %d ms\n", msg, NPASSES, delta);
}