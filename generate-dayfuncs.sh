#!/bin/bash
days=`ls -1 day*.go | egrep '^day[0-9]+.go$' | sed -e 's/^day//' -e 's/\.go//'`
# var dayFuncs = map[int]year.Day{
# 	1: &Day1{},
# }
outfile=dayfuncs.go
echo "// Auto-generated from created files" > $outfile

echo -e `head -1 aoc.go` >> $outfile
echo "" >> $outfile
echo -e "import \"github.com/jbert/aoc/year\"\n" >> $outfile
echo -e "var dayFuncs = map[int]year.Day{" >> $outfile
for day in $days
do
  echo -e "\t$day: &Day$day{}," >> $outfile
done 
echo "}" >> $outfile
