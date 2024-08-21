# golang_pdf_eof_fix
Golang pdf EOF file fix

This package was created to serve as a tool repository for fixing EOF on PDF files.

FixEofOnPdfFile method will search for the first '%%EOF' match and subtract the bytes further from this point and
generate a new file.

** WARNING ** This should be used with caution as it will replace the original PDF file with a new one with EOF mark adjusted.
Be sure to have a backup of original files.