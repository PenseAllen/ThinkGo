#!/usr/bin/env python3

import re
import sys

RE_TERM = re.compile(r'\*([a-z].+?)\*', re.DOTALL)
RE_ENTRY = re.compile(r'^(.+)::', re.MULTILINE)

def find_terms(text):
    terms = set()
    duplicates = set()
    text = text.split('=== Glossary')[0]
    for term in RE_TERM.findall(text):
        if 'all valid tokens has' in term:
            continue
        if '\n' in term:
            term = ' '.join(term.split())
        if term.endswith('s'):
            term = term[:-1]
        if term in terms:
            duplicates.add(term)
        else:
            terms.add(term)
    return terms, duplicates

def find_entries(text):
    entries = set()
    glossary = text.split('=== Glossary')[1]
    for entry in RE_ENTRY.findall(glossary):
        entries.add(entry)
    return entries

def main(text_name):
    with open(text_name) as fp:
        text = fp.read()

    terms, duplicates = find_terms(text)

    print(','.join(t for t in sorted(terms)))

    print()
    if duplicates:
        print('DUPLICATES:')
        for term in sorted(duplicates):
            print('\t' + term)

    entries = find_entries(text)
    print('ORPHAN TERMS:')
    for term in sorted(terms-entries):
        print('\t' + term)

    print('ORPHAN ENTRIES:')
    for term in sorted(entries-terms):
        print('\t' + term)





if __name__ == '__main__':
    main(sys.argv[1])
