//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_FACTORY_H
#define CODE_CPP_FACTORY_H

#include "ConcreteProduct.h"

class Factory
{
public:
    virtual Car* CreateCar() = 0;
};

#endif //CODE_CPP_FACTORY_H