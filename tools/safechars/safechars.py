import unicodedata
for octet in range(0x80, 0x100):
    b = bytes([octet])
    from437 = b.decode('cp437')
    from850 = b.decode('cp850')
    if from437 == from850:
        try:
            name = unicodedata.name(from437)
        except ValueError:
            continue
        if 'LATIN' in name:
            continue

        #print('{octet:02x} {from437} {name}'.format(**locals()))
        print(from437, end=' ')
