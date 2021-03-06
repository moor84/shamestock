# Shamestock
Command-line toolkit to help migrating out of Shutterstock.

Shamestock can help in moving a raster or vector stock portfolio out of Shutterstock,
prepare and upload it to other stock illustration platforms.

Requires *exiv2* and *imagemagick* console utils to be present in the system.
They can be installed, for example, with Homebrew:
```
brew install exiv2
brew install imagemagick
```

## Preparing a batch
Suppose you have a batch of .eps vector files stored in batches/batch1.
Tipically, a process of preparing a batch of vectors for uploading consists of the following commands:

```
shamestock prepare template batches/batch1/
```

Creates a template for a CSV file, batches/batch1/urls.csv, e.g.:
```
1234.eps,<place for shutterstock url>
1235.eps,<place for shutterstock url>
...
```

Once you fill in the urls in the csv file, you can proceed further
(the rest of the steps are fully automated).
```
shamestock meta scrapecsv batches/batch1/urls.csv
```
Scrapes titles and keywords from Shutterstock and stores them, alongside the file names,
to batches/batch1/attrs.csv
```
shamestock preview generate batches/batch1
```
Generates .jpg previews for all the .eps files in batches/batch1.
```
shamestock meta csv batches/batch1/attrs.csv
```
Writes titles and keywords from batches/batch1/attrs.csv to correspoding jpeg files in
the batches/batch1/ directory.
```
shamestock prepare zip batches/batch1
```
Create zip archives required, for example, for Adobe Stock, and store them in
the batches/batch1/zip directory.
