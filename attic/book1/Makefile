LATEX = latex

DVIPS = dvips

PDFFLAGS = -dCompatibilityLevel=1.4 -dPDFSETTINGS=/prepress \
           -dCompressPages=true -dUseFlateCompression=true  \
           -dEmbedAllFonts=true -dSubsetFonts=true -dMaxSubsetPct=100


%.dvi: %.tex
	$(LATEX) $<

%.ps: %.dvi
	$(DVIPS) -o $@ $<

%.pdf: %.ps
	ps2pdf $(PDFFLAGS) $<

all:	book.tex
	pdflatex book
	makeindex book.idx
	pdflatex book
	mv book.pdf thinkgo.pdf

hevea:	book.tex header.html footer.html
	# replace the pdfs with eps
	sed s/.pdf/.eps/g book.tex > thinkgo.tex
	latex thinkgo
	rm -rf html
	mkdir html
	hevea -fix -O -e latexonly htmlonly thinkgo
# the following greps are a kludge to prevent imagen from seeing
# the definitions in latexonly, and to avoid headers on the images
	grep -v latexonly thinkgo.image.tex > a; mv a thinkgo.image.tex
	sed s/\\\\usepackage{fancyhdr}// < thinkgo.image.tex > a; mv a thinkgo.image.tex
	imagen -png thinkgo
	hacha thinkgo.html
	cp up.png next.png back.png html
	mv index.html thinkgo.css thinkgo*.html thinkgo*.png *motif.gif html

DEST = /home/downey/public_html/greent/thinkgo

epub:
	cd html; ebook-convert index.html thinkgo.epub

distrib:
	rm -rf dist
	mkdir dist dist/tex dist/tex/figs
	rsync -a thinkgo.pdf html dist
	rsync -a dist/* $(DEST)
	chmod -R o+r $(DEST)/*
	cd $(DEST)/..; sh back

# UPDATE THE PATHS BELOW BEFORE RUNNING PLASTEX

plastex:
	# Before running plastex, we need the current directory in PYTHONPATH
	# export PYTHONPATH=$PYTHONPATH:.
	python Filist.py book.tex > book.plastex
	rm -rf /home/downey/thinkgo/trunk/book
	plastex --renderer=DocBook --theme=book --image-resolution=300 --filename=book.xml book.plastex
	rm -rf /home/downey/thinkgo/trunk/book/.svn

plastest:
	# Before running plastex, we need the current directory in PYTHONPATH
	# export PYTHONPATH=$PYTHONPATH:.
	python Filist.py test.tex > test.plastex
	rm -rf /home/downey/thinkgo/trunk/test
	plastex --renderer=DocBook --theme=test --filename=test.xml test.plastex
	rm -rf /home/downey/thinkgo/trunk/test/.svn

xxe:
	xmlcopyeditor ~/Thinkgo/book/book/book.xml &

lint:
	xmllint -noout book/book.xml

OREILLY = atlas

oreilly:
	rsync -a book.tex $(OREILLY)
	rsync -a book/ $(OREILLY)
	rsync -a figs/* $(OREILLY)/figs
	cd $(OREILLY); git add .
	cd $(OREILLY); git commit -m "Automated check in."
	cd $(OREILLY); git push

clean:
	rm -f *~ *.aux *.log *.dvi *.idx *.ilg *.ind *.toc
