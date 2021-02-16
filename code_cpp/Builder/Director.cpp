//
// Created by ZXY on 2021/2/16.
//

#include "Director.h"

void Director::CreateComputer(ComputerBuilder *builder) {
    builder->BuildCpu();
    builder->BuildMainboard();
    builder->BuildRam();
    builder->BuildVideoCard();
}