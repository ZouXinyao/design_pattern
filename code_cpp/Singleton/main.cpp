//
// Created by ZXY on 2021/2/16.
//

#include "Singleton.h"

int main()
{
    Singleton* single = Singleton::GetInstance();
    single->ProveSuccSingleton();

    return 0;
}
