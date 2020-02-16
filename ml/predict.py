import csv
import pickle
import numpy as np
from xgboost import XGBClassifier
from sklearn.model_selection import StratifiedKFold, train_test_split
from sklearn.model_selection import RandomizedSearchCV, GridSearchCV
from sklearn.metrics import confusion_matrix, f1_score
from sklearn.externals import joblib
from scipy import stats
import matplotlib.pyplot as plt

if __name__ == "__main__":
    # data = []
    # csvfile = open("modifiedData2.csv", newline='')
    # reader = csv.reader(csvfile)
    # for line in reader:
    #     data.append(line)

    # csvfile.close()

    # data = [[float(n) for n in line] for line in data]

    # one = [line[:-1] for line in data if line[-1]==1.0]
    # zero = [line[:-1] for line in data if line[-1]==0.0]

    # x = one + zero
    # x = np.array(x)
    # y = [1 for x in one] + [0 for x in zero]
    # y = np.array(y)

    # x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.05, random_state=10, stratify=y)
    clf = joblib.load("model.joblib")
    # clf = pickle.load(open('model.joblib', 'rb'))
    ##X,Y,MONTH,FFMC,DMC,DC,ISI,TEMP,RH,WIND,RAIN
    input = [5,5,6,12,90,70,650,10,50,0,0]
    y_pred = clf.predict(input)
    print(y_pred)
    print(clf.best_score_)