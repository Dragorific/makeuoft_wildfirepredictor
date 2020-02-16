import csv
import numpy as np
import pandas as pd
import seaborn
import matplotlib.pyplot as plt

months = ["jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"]

def numMonth(a):
    return [months.index(a)]

if __name__ == "__main__":

    data = []
    csvfile = open("forestfires.csv", newline='')
    outfile = open('modifiedData2.csv', 'w', newline='')
    reader = csv.reader(csvfile)
    writer = csv.writer(outfile)
    print(reader.__next__())
    for line in reader:
        if (float(line[-1]) > 0):
            temp  = line[0:2]+numMonth(line[2])+line[4:-1]
            temp.append("1")
            writer.writerow(temp)
        else:
            writer.writerow(line[0:2]+numMonth(line[2])+line[4:])

    csvfile.close()
    outfile.close()