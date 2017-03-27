import wcwidth

for x in range(0, 0x10FFFF):
    print("%04x %s" % (x, wcwidth.wcwidth(chr(x))))
