//
// Created by ZXY on 2021/2/16.
//

#include "Singleton.h"
#include <iostream>

Singleton* Singleton::m_Singleton = nullptr;

Singleton* Singleton::GetInstance() {
    if (m_Singleton == nullptr) {
        m_Singleton = new Singleton();
    }
    return m_Singleton;
}

void Singleton::ProveSuccSingleton() {
    std::cout << "Creating Singleton success" << std::endl;
}

Singleton::Singleton() = default;
