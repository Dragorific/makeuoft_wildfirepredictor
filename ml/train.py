import csv
import numpy as np
from xgboost import XGBClassifier
from sklearn.model_selection import StratifiedKFold, train_test_split
from sklearn.model_selection import RandomizedSearchCV, GridSearchCV
from sklearn.metrics import confusion_matrix, f1_score
from sklearn.externals import joblib
from scipy import stats
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

    x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.05, random_state=10, stratify=y)
    num_class = len(list(set(y_train)))

    clf_xgb = XGBClassifier(objective='multi:softmax', num_class=num_class, n_jobs=-1, nthread=-1)
    param_dist = {'n_estimators': np.array(range(25, 500, 25)),
                'learning_rate': stats.uniform(0.1, 0.59),#0.01+0.59
                'subsample': stats.uniform(0.4, 0.6),
                'max_depth': np.array([3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]),
                'colsample_bytree': stats.uniform(0.4, 0.3),
                'min_child_weight': np.array([1, 2, 3, 4, 5, 6])
                }
    n_folds = 3
    kfold = StratifiedKFold(n_splits=n_folds, shuffle=False, random_state=10)
    clf = RandomizedSearchCV(clf_xgb,
                            param_distributions=param_dist,
                            cv=kfold,
                            n_iter=100,
                            scoring='f1',
                            error_score=1,
                            verbose=0,
                            refit=True,
                            n_jobs=-1)
    clf.fit(x_train, y_train)
    joblib.dump(clf, 'model.joblib')
    # clf_xgb.fit(x_train, y_train)
    # clf_xgb.fit(x_train, y_train)
    y_pred = clf.predict(x_test)
    random_model, random_params = clf.best_estimator_, clf.best_params_
    # random_model, random_params = clf_xgb.best_estimator_, clf_xgb.best_params_


    # n_estimators = [int(x) for x in np.linspace(int(random_params['n_estimators'] - 5), int(random_params['n_estimators'] + 5), num=3) if x > 0]
    # learning_rate = [float(x) for x in np.linspace(float(random_params['learning_rate'] - 0.02), float(random_params['learning_rate'] + 0.02), num=5) if x > 0]
    # subsample = [float(x) for x in np.linspace(float(random_params['subsample'] - 0.02), float(random_params['subsample'] + 0.02), num=5) if x > 0]
    # colsample_bytree = [float(x) for x in np.linspace(float(random_params['colsample_bytree'] - 0.02), float(random_params['colsample_bytree'] + 0.02), num=5) if x > 0]
    # max_depth = [int(random_params['max_depth'])]
    # min_child_weight = [int(random_params['min_child_weight'])]

    # param_dist = {'n_estimators': n_estimators,
    #             'learning_rate': learning_rate,
    #             'subsample': subsample,
    #             'max_depth': max_depth,
    #             'colsample_bytree': colsample_bytree,
    #             'min_child_weight': min_child_weight
    #             }

    # clf = GridSearchCV(random_model,
    #                 param_grid=param_dist,
    #                 cv=kfold,
    #                 scoring='f1_weighted',
    #                 error_score=0,
    #                 verbose=0,
    #                 refit=True,
    #                 n_jobs=-1)

    # # Let's fit our model.
    # clf.fit(x_train, y_train)

    # # Finally, we'll extract our best model and parameters, which we will use for testing.
    # grid_model, grid_params = clf.best_estimator_, clf.best_params_

    # print(grid_model)
    # print(grid_params)
    print(random_model)
    print(random_params)
    print(clf.best_score_)