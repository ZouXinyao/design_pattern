//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_FACTORY_H
#define CODE_CPP_FACTORY_H

#include "ConcreteProductCar.h"
#include "ConcreteProductShip.h"

class Factory
{
public:
    enum FACTORY_TYPE {
        BENZ_FACTORY,
        BMW_FACTORY,
    };

    virtual Car* CreateCar() = 0;
    virtual Ship* CreateShip() = 0;
    static Factory* CreateFactory(FACTORY_TYPE factoryType);
};

#endif //CODE_CPP_FACTORY_H