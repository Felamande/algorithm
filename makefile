DGB = -g
CC = gcc
TGT = 

OBJS = $(TGT).o stack.o time.o



$(TGT).exe:$(OBJS)
	$(CC) $(DBG) -o $(TGT).exe $(OBJS)

$(OBJS):%.o:%.c
	$(CC) -c $(DGB) $< -o $@  

clean:
	-del *.o $(TGT)


.PHONY: clean