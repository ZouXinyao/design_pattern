//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_CONCRETEFACTORY_H
#define CODE_CPP_CONCRETEFACTORY_H

#include "Factory.h"
#include "ConcreteProduct.h"

class BenzFactory : public Factory
{
public:
    Car* CreateCar() override;
};

// ±¦Âí¹¤³§
class BmwFactory : public Factory
{
public:
    Car* CreateCar() override;
};


#endif //CODE_CPP_CONCRETEFACTORY_H
