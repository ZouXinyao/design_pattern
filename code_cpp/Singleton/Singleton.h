//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_SINGLETON_H
#define CODE_CPP_SINGLETON_H

class Singleton {
private:
    static Singleton *m_Singleton;
    Singleton();

public:
    static Singleton* GetInstance();
    void ProveSuccSingleton();
};

#endif //CODE_CPP_SINGLETON_H
