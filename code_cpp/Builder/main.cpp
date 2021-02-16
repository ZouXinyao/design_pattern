//
// Created by ZXY on 2021/2/16.
//

#include "ConcreteBuilder.h"
#include "Director.h"
#include <string>

int main()
{
    auto *pDirecror = new Director();
    auto *pLenovoBuilder = new LenovoBuilder();
    auto *pDellBuilder = new DellBuilder();

    // ��װ ThinkPad��Yoga
    pDirecror->CreateComputer(pLenovoBuilder);
    pDirecror->CreateComputer(pDellBuilder);

    // ��ȡ��װ��ĵ���
    Computer *pLenovoComputer = pLenovoBuilder->GetComputerResult();
    Computer *pDellComputer = pDellBuilder->GetComputerResult();

    // �������
    cout << "-----Lenovo-----" << endl;
    cout << "CPU: " << pLenovoComputer->GetCPU() << endl;
    cout << "Mainboard: " << pLenovoComputer->GetMainboard() << endl;
    cout << "Ram: " << pLenovoComputer->GetRam() << endl;
    cout << "VideoCard: " << pLenovoComputer->GetVideoCard() << endl;

    cout << "-----Dell-----" << endl;
    cout << "CPU: " << pDellComputer->GetCPU() << endl;
    cout << "Mainboard: " << pDellComputer->GetMainboard() << endl;
    cout << "Ram: " << pDellComputer->GetRam() << endl;
    cout << "VideoCard: " << pDellComputer->GetVideoCard() << endl;

    delete(pLenovoBuilder);
    delete(pDellBuilder);
    delete(pDirecror);

    return 0;
}
