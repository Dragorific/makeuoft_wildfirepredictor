import csv
import numpy as np
import pandas as pd
import seaborn
import matplotlib.pyplot as plt

data = []
csvfile = open("forestfires.csv", newline='')
outfile = open('modifiedData.csv', 'w', newline='')
reader = csv.reader(csvfile)
writer = csv.writer(outfile)
print(reader.__next__())
for line in reader:
    if (float(line[-1]) > 0):
        temp  = line[5:-1]
        temp.append("1")
        writer.writerow(temp)
        data.append(line[5:-1].append('1'))
    else:
        writer.writerow(line[5:])
        data.append(line[5:])

csvfile.close()
outfile.close()