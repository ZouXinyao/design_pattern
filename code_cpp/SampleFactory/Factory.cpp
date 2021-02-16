//
// Created by ZXY on 2021/2/16.
//

#include "Factory.h"

Car *Factory::CreateCar(Factory::AUTOMOBLIE_BRAND brand) {
    Car *pCar = nullptr;

    // �򵥹���ʹ��if-else�߼���֧�ж�
    if(brand == AUTOMOBLIE_BRAND::BENZ_BRAND) pCar = new Benz();
    else if(brand == AUTOMOBLIE_BRAND::BMW_BRAND) pCar = new BMW();

    return pCar;
}
