import csv
import numpy as np
import pandas as pd
import seaborn
import matplotlib.pyplot as plt

data = []
csvfile = open("modifiedData.csv", newline='')
reader = csv.reader(csvfile)
for line in reader:
    data.append(line)

csvfile.close()

data = [[float(n) for n in line] for line in data]

one = [line[:-1] for line in data if line[-1]==1.0]
zero = [line[:-1] for line in data if line[-1]==0.0]
print(len(one))
print(len(zero))

# df = pd.DataFrame(one)

# df_corr = df.corr()

# plt.figure(figsize=(10,10))
# seaborn.heatmap(df_corr, cmap="Reds")
# seaborn.set(font_scale=2,style='white')

# plt.title('Heatmap correlation')
# plt.show()
