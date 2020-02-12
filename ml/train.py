import csv
import numpy as np
from xgboost import XGBClassifier
from sklearn.model_selection import StratifiedKFold, train_test_split
from sklearn.model_selection import RandomizedSearchCV, GridSearchCV
from sklearn.metrics import confusion_matrix, f1_score
from scipy import stats
import matplotlib.pyplot as plt

if __name__ == "__main__":

    data = []
    csvfile = open("modifiedData.csv", newline='')
    reader = csv.reader(csvfile)
    for line in reader:
        data.append(line)

    csvfile.close()

    data = [[float(n) for n in line] for line in data]

    one = [line[:-1] for line in data if line[-1]==1.0]
    zero = [line[:-1] for line in data if line[-1]==0.0]

    x = one + zero
    y = [1 for x in one] + [0 for x in zero]
    print(x)
    exit(0)
    x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.2, random_state=10, stratify=y)

    num_class = len(list(set(y_train)))
    clf_xgb = XGBClassifier(objective='multi:softmax', num_class=num_class, n_jobs=-1, nthread=-1)
    param_dist = {'n_estimators': np.array(range(100, 1000, 100)),
                'learning_rate': stats.uniform(0.01, 0.59),#0.01+0.59
                'subsample': stats.uniform(0.3, 0.6),
                'max_depth': np.array([3, 4, 5, 6, 7, 8, 9]),
                'colsample_bytree': stats.uniform(0.5, 0.3),
                'min_child_weight': np.array([1, 2, 3, 4])
                }
    n_folds = 3
    kfold = StratifiedKFold(n_splits=n_folds, shuffle=False, random_state=10)
    clf = RandomizedSearchCV(clf_xgb,
                            param_distributions=param_dist,
                            cv=kfold,
                            n_iter=100,
                            scoring='f1_weighted',
                            error_score=0,
                            verbose=0,
                            refit=True,
                            n_jobs=-1)
    clf.fit(x_train, y_train)
    random_model, random_params = clf.best_estimator_, clf.best_params_

    print(random_model, '\n')
    print(random_params, '\n')
    print(clf.best_score_)