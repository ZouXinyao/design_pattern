//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_FACTORY_H
#define CODE_CPP_FACTORY_H

#include "ConcreteProduct.h"

class Factory
{
public:
    enum AUTOMOBLIE_BRAND {
        BENZ_BRAND,
        BMW_BRAND
    };

    Car *CreateCar(AUTOMOBLIE_BRAND brand);
};

#endif //CODE_CPP_FACTORY_H