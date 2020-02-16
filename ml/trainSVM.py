import numpy as np
import pandas as pd
import csv
import seaborn
from sklearn import svm
from sklearn.model_selection import train_test_split
from sklearn.metrics import confusion_matrix, classification_report
import matplotlib.pyplot as plt

if __name__ == "__main__":

    data = []
    csvfile = open("modifiedData2.csv", newline='')
    reader = csv.reader(csvfile)
    for line in reader:
        data.append(line)

    csvfile.close()

    data = [[float(n) for n in line] for line in data]

    one = [line[:-1] for line in data if line[-1]==1.0]
    zero = [line[:-1] for line in data if line[-1]==0.0]

    x = one + zero
    x = np.array(x)
    y = [1 for x in one] + [0 for x in zero]
    y = np.array(y)

    x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.1, random_state=10, stratify=y)

    clf = svm.SVC()
    clf.fit(x_train, y_train)

    y_pred = list(clf.predict(x_test))
    

    tn, fp, fn, tp = confusion_matrix(y_test, y_pred).ravel()
    print(tn, fp, fn, tp)
    results = [[tp, fn], [fp, tn]]

    x_labels = ['Predicted True', 'Predicted False']
    y_labels = ['Actual True', 'Actual False']

    fig, ax = plt.subplots()
    im = ax.imshow(results, cmap='Reds')

    ax.set_xticks(np.arange(len(x_labels)))
    ax.set_yticks(np.arange(len(y_labels)))

    ax.set_xticklabels(x_labels)
    ax.set_yticklabels(y_labels)

    plt.setp(ax.get_xticklabels(), rotation=45, ha="right",
            rotation_mode="anchor")

    for i in range(len(x_labels)):
        for j in range(len(y_labels)):
            text = ax.text(j, i, results[i][j],
                        ha="center", va="center", color="black")

    ax.set_ylim(len(results)-0.5, -0.5)
    ax.set_title("Confusion Matrix")
    fig.tight_layout()
    plt.show()