package Utils;

import java.util.List;

/**
 * Created by zl on 2017/4/30.
 */

public class gd11x5 {
    /**
     * 排列数
     * @param n
     * @param m
     * @return
     */
    public int PlNum(int n,int m) {
        int x = 1;int y = 1;
        for (int i=n;i>0;i--) {
            x *= i;
        }

        if(n-m ==0 || n-m == 1) {
            y = 1;
        }else {
            for (int i= n-m;i>0;i--) {
                y *= i;
            }
        }
        return x/y;
    }

    /**
     * 组合数 n选m
     * @return
     */
    public int ZHnum(int n,int m) {
        int x=1;int y=1;int z =1;
        if(n<m) {
            return 0;
        }
        if(n == m) {
            x= 1;
            y = 1;
            z = 1;
        }else {
            for (int i=1;i<=n;i++) {
                x *= i;
            }
            for (int j=1;j<= m;j++) {
                y *= j;
            }
            for (int k =1;k<=n-m;k++) {
                z *= k;
            }
        }
        return x/(y*z);
    }
    /**
     * 返回任选二普通投注注数
     * @param mList
     * @return
     */
    public int NormalRX2(List<String> mList) {
        int Zs = 0;

        Zs = ZHnum(mList.size(),2);
        return Zs;
    }
    /**
     * 返回任选三普通投注注数
     * @param mList
     * @return
     */
    public int NormalRX3(List<String> mList) {
        int Zs = 0;

        Zs = ZHnum(mList.size(),3);
        return Zs;
    }
    /**
     * 返回任选四普通投注注数
     * @param mList
     * @return
     */
    public int NormalRX4(List<String> mList) {
        int Zs = 0;

        Zs = ZHnum(mList.size(),4);
        return Zs;
    }

    /**
     * 返回任选5普通投注注数
     * @param mList
     * @return
     */
    public int NormalRX5(List<String> mList) {
        int Zs = 0;

        Zs = ZHnum(mList.size(),5);
        return Zs;
    }
    /**
     * 返回任选6普通投注注数
     * @param mList
     * @return
     */
    public int NormalRX6(List<String> mList) {
        int Zs = 0;

        Zs = ZHnum(mList.size(),6);
        return Zs;
    }
    /**
     * 返回任选7普通投注注数
     * @param mList
     * @return
     */
    public int NormalRX7(List<String> mList) {
        int Zs = 0;

        Zs = ZHnum(mList.size(),7);
        return Zs;
    }
    /**
     * 返回任选8普通投注注数
     * @param mList
     * @return
     */
    public int NormalRX8(List<String> mList) {
        int Zs = 0;

        Zs = ZHnum(mList.size(),8);
        return Zs;
    }

    /**
     * 返回任前一普通投注注数
     * @param mList
     * @return
     */
    public int NormalQy(List<String> mList) {
        return mList.size();
    }
    /**
     * 返回任前二普通投注注数
     * @param mList
     * @return
     */
    public int NormalQe(List<String> mList,List<String> twoLineList) {
        int ZS = 1;
        int com = 0;
        int amonutSize = mList.size()+twoLineList.size();

        if(mList.size() <1) {
            return 0;
        }
        if(twoLineList.size()<1) {
            return 0;
        }
        for (int i =0;i<mList.size();i++) {
            for (int j=0;j<twoLineList.size();j++) {
                if(mList.get(i).equals(twoLineList.get(j))) {
                    com += 1;
                }
            }
        }
        if(amonutSize - com <2){
            return 0;
        }
        ZS = mList.size() * twoLineList.size() - com;
        return ZS;
    }

    /**
     * 拖胆任选二注数 前二
     * @param oneLine 胆码
     * @param twoLine 拖码
     * @return
     */
    public int TdRx2(List<String> oneLine,List<String> twoLine) {
        int ZS = 0;
        if(oneLine.size() == 0) {
            return 0;
        }
        if (twoLine.size() == 0) {
            return 0;
        }
        ZS = twoLine.size();
        return ZS;
    }

    /**
     * 拖胆任选3注数 前三
     * @param oneLine 胆码
     * @param twoLine 拖码
     * @return
     */
    public int TdRx3(List<String> oneLine,List<String> twoLine) {
        int ZS = 0;
        if(oneLine.size() == 0) {
            return 0;
        }
        if (twoLine.size() == 0) {
            return 0;
        }
        ZS = ZHnum(twoLine.size(),3-oneLine.size());
        return ZS;
    }
    /**
     * 拖胆任选4注数
     * @param oneLine 胆码
     * @param twoLine 拖码
     * @return
     */
    public int TdRx4(List<String> oneLine,List<String> twoLine) {
        int ZS = 0;
        if(oneLine.size() == 0) {
            return 0;
        }
        if (twoLine.size() == 0) {
            return 0;
        }
        ZS = ZHnum(twoLine.size(),4-oneLine.size());
        return ZS;
    }
    /**
     * 拖胆任选5注数
     * @param oneLine 胆码
     * @param twoLine 拖码
     * @return
     */
    public int TdRx5(List<String> oneLine,List<String> twoLine) {
        int ZS = 0;
        if(oneLine.size() == 0) {
            return 0;
        }
        if (twoLine.size() == 0) {
            return 0;
        }
        ZS = ZHnum(twoLine.size(),5-oneLine.size());
        return ZS;
    }
    /**
     * 拖胆任选6注数
     * @param oneLine 胆码
     * @param twoLine 拖码
     * @return
     */
    public int TdRx6(List<String> oneLine,List<String> twoLine) {
        int ZS = 0;
        if(oneLine.size() == 0) {
            return 0;
        }
        if (twoLine.size() == 0) {
            return 0;
        }
        ZS = ZHnum(twoLine.size(),6-oneLine.size());
        return ZS;
    }
    /**
     * 拖胆任选7注数
     * @param oneLine 胆码
     * @param twoLine 拖码
     * @return
     */
    public int TdRx7(List<String> oneLine,List<String> twoLine) {
        int ZS = 0;

        if(oneLine.size() == 0) {
            return 0;
        }
        if (twoLine.size() == 0) {
            return 0;
        }
        ZS = ZHnum(twoLine.size(),7-oneLine.size());
        return ZS;
    }
    /**
     * 拖胆任选8注数
     * @param oneLine 胆码
     * @param twoLine 拖码
     * @return
     */
    public int TdRx8(List<String> oneLine,List<String> twoLine) {
        int ZS = 0;

        if(oneLine.size() == 0) {
            return 0;
        }
        if (twoLine.size() == 0) {
            return 0;
        }
        ZS = ZHnum(twoLine.size(),8-oneLine.size());
        return ZS;
    }
    /**
     * 前三组选投注注数
     * @param mList
     * @return
     */
    public int NormalQs(List<String> mList) {
        int Zs = 0;
        Zs =  ZHnum(mList.size(),2);
        return Zs;
    }
    /**
     * 前二组选投注注数
     * @param mList
     * @return
     */
    public int NormalQE(List<String> mList) {
        int Zs = 0;

        Zs = ZHnum(mList.size(),3);
        return Zs;
    }
}
