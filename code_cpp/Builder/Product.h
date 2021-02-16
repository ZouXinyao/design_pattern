//
// Created by ZXY on 2021/2/16.
//

#ifndef CODE_CPP_PRODECT_H
#define CODE_CPP_PRODECT_H

#include <iostream>

using namespace std;

class Computer
{
public:
    void SetCpu(string cpu)             { m_Cpu       = cpu; }
    void SetMainboard(string mainboard) { m_Mainboard = mainboard; }
    void SetRam(string ram)             { m_Ram       = ram; }
    void SetVideoCard(string videoCard)  { m_VideoCard = videoCard; }

    string GetCPU()        { return m_Cpu; }
    string GetMainboard()  { return m_Mainboard; }
    string GetRam()        { return m_Ram; }
    string GetVideoCard()  { return m_VideoCard; }

private:
    string m_Cpu;  // CPU
    string m_Mainboard;  // Ö÷°å
    string m_Ram;  // ÄÚ´æ
    string m_VideoCard;  // ÏÔ¿¨
};

#endif //CODE_CPP_PRODECT_H
