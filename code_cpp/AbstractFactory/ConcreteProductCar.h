//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_CONCRETEPRODUCT_H
#define CODE_CPP_CONCRETEPRODUCT_H

#include "ProductCar.h"

class BenzCar : public Car
{
public:
    std::string Name() override;
};

class BMWCar : public Car
{
public:
    std::string Name() override;
};


#endif //CODE_CPP_CONCRETEPRODUCT_H
