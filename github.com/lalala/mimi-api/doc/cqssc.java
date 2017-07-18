package Utils;

import java.util.ArrayList;
import java.util.List;

import LottoryUtils.LottoryBalls;

/**
 * Created by zl on 2017/4/30.
 */

public class Cqssc {
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
    //五星直选通选

    public int setFiveZT(List<String> WwNos,List<String> QwNos,List<String> BwNos,List<String> SwNos,List<String> GwNos) {
        int ZS = 1;
        if(WwNos.size() <1) {
            return 0;
        }
        if(QwNos.size() <1) {
            return 0;
        }
        if(BwNos.size() <1) {
            return 0;
        }
        if(SwNos.size() <1) {
            return 0;
        }
        if(GwNos.size() <1) {
            return 0;
        }
        ZS = WwNos.size() * QwNos.size()*BwNos.size()*SwNos.size()*GwNos.size();
        return  ZS;
    }
    //五星组选120  ZHnum（WwNos.size(),5）;



    /**
     * 设置五星组选60注数 (i-k)cj3+k*C(j-1)3//i二从个数,j三单个数，K重发个数
     com 重复数
     */
    public int setFive60(List<String> WwNos,List<String> QwNos) {
        int com = 0;
        int amonutSize = WwNos.size()+QwNos.size();
        int Zs = 1;
        if(WwNos.size()<1) {
            return 0;
        }
        if(QwNos.size()<3) {
            return 0;
        }
        for (int i =0;i<WwNos.size();i++) {
            for (int j=0;j<QwNos.size();j++) {
                if(WwNos.get(i).equals(QwNos.get(j))) {
                    com += 1;
                }
            }
        }
        if(amonutSize - com <4) {
            return 0;
        }

        int x = (WwNos.size()-com)* ZHnum(QwNos.size(),3);
        int y = com*ZHnum(QwNos.size()-1,3);
        Zs = x+y;
        return Zs;
    }

    /**
     * 设置五星组选30
     * i 单号个数，
     n 二重号个数
     k 重复个数
     (i-k)*Cn2  + (k)*C(n-1)2
     */
    public int setFive30(List<String> WwNos,List<String> QwNos) {
        int com = 0;
        int amonutSize = WwNos.size()+QwNos.size();
        int Zs = 1;
        if(WwNos.size() <2) {
            return 0;
        }
        if(QwNos.size() <1) {
            return 0;
        }
        for (int i =0;i<WwNos.size();i++) {
            for (int j=0;j<QwNos.size();j++) {
                if(WwNos.get(i).equals(QwNos.get(j))) {
                    com += 1;
                }
            }
        }
        if (amonutSize -com<3) {
            return 0;
        }
        int x = (QwNos.size()-com)*ZHnum(WwNos.size(),2);
        int y = com*ZHnum(WwNos.size()-1,2);
        Zs = x+y;
        return Zs;
    }


    /**
     * 设置五星组20注数
     * i 三重号个数，
     n 单号个数
     k 重复个数
     (i-k)*Cn2  + (k)*C(n-1)2
     */
    public int setFive20(List<String> WwNos,List<String> QwNos) {
        int com = 0;
        int amonutSize = WwNos.size()+QwNos.size();
        int Zs = 1;
        if(WwNos.size()<1) {
            return 0;
        }
        if(QwNos.size()<2) {
            return 0;
        }
        for (int i =0;i<WwNos.size();i++) {
            for (int j=0;j<QwNos.size();j++) {
                if(WwNos.get(i).equals(QwNos.get(j))) {
                    com += 1;
                }
            }
        }
        if (amonutSize -com<3) {
            return 0;
        }
        int x = (WwNos.size()-com)*ZHnum(QwNos.size(),2);
        int y = com*ZHnum(QwNos.size()-1,2);
        Zs = x+y;

        return Zs;
    }

    /**
     * 设置五星组10和组5注数
     * i 重号个数，
     n 单号个数
     k 重复个数
     (i-k)*n  + (k)*（n-1）
     */
    public int setTenAndFive(List<String> WwNos,List<String> QwNos) {
        int com = 0;
        int amonutSize = WwNos.size()+QwNos.size();
        int Zs = 1;
        if(WwNos.size() <1) {
            return 0;
        }
        if(QwNos.size()<1) {
            return 0;
        }
        for (int i =0;i<WwNos.size();i++) {
            for (int j=0;j<QwNos.size();j++) {
                if(WwNos.get(i).equals(QwNos.get(j))) {
                    com += 1;
                }
            }
        }
        if(amonutSize -com <2) {
            return 0;
        }
        int x = (WwNos.size()-com)*QwNos.size();
        int y = com*(QwNos.size()-1);
        Zs = x+y;
        return Zs;
    }

    /**
     * 前四后四直选注数
     */
    public int setFourZX(List<String> QwNos,List<String> BwNos,List<String> SwNos,List<String> GwNos) {
        int ZS = 1;
        if(QwNos.size() <1) {
            return 0;
        }
        if(BwNos.size()<1) {
            return 0;
        }
        if(SwNos.size()<1) {
            return 0;
        }
        if(GwNos.size() <1) {
            return 0;
        }
        ZS = QwNos.size()*BwNos.size()*GwNos.size()*SwNos.size();
        return ZS;
    }


    /**
     * 四星24注数
     * @return 注数
     */
    public int setFour24(List<String> WwNos) {
        int ZS = 1;

        if(WwNos.size()<4) {
            return 0;
        }
        ZS = ZHnum(WwNos.size(),4);
        return ZS;
    }


    /**
     * 四星12注数
     * @return 注数
     * i 二重号个数，
    n 单号个数
    k 重复个数
    (i-k)*Cn2  + (k)*C(n-1)2  com 重复数

     */
    public int setFour12(List<String> WwNos,List<String> QwNos) {
        int ZS = 1;
        int com = 0;
        int amonutSize = WwNos.size()+QwNos.size();
        if(WwNos.size()<1) {
            return 0;
        }
        if(QwNos.size() <2) {
            return 0;
        }
        for (int i =0;i<WwNos.size();i++) {
            for (int j=0;j<QwNos.size();j++) {
                if(WwNos.get(i).equals(QwNos.get(j))) {
                    com += 1;
                }
            }
        }
        if(amonutSize - com <3) {
            return 0;
        }
        int x = (WwNos.size()-com)*ZHnum(QwNos.size(),2);
        int y = com * ZHnum(QwNos.size()-1,2);
        ZS = x+y;
        return ZS;
    }

    /**
     * 四星组六、组四注数
     * i 二重号（第一行）个数，
     n 二重号（第二行）个数
     k 重复个数
     (i-k)*n  + (k)*（n-1）
     */
    public int setFourSix(List<String> WwNos,List<String> QwNos) {
        int ZS = 1;
        int com = 0;
        int amonutSize = WwNos.size()+QwNos.size();
        if(WwNos.size() <1) {
            return 0;
        }
        if(QwNos.size()<1) {
            return 0;
        }
        for (int i =0;i<WwNos.size();i++) {
            for (int j=0;j<QwNos.size();j++) {
                if(WwNos.get(i).equals(QwNos.get(j))) {
                    com += 1;
                }
            }
        }
        if(amonutSize - com <2){
            return 0;
        }

        int x= (WwNos.size()-com)*QwNos.size();
        int y = com*(QwNos.size()-1);
        ZS = x*y;
        return ZS;
    }
    /**
     * 前三 z中三，后三直选注数
     * @return 注数
     */
    public int setFrtTreZX(List<String> WwNos,List<String> QwNos,List<String> BwNos) {
        int ZS=1;
        if(WwNos.size() <1) {
            return 0;
        }
        if(QwNos.size() <1) {
            return 0;
        }
        if(BwNos.size()<1) {
            return 0;
        }
        ZS = WwNos.size()*QwNos.size()*BwNos.size();
        return ZS;
    }

    /**
     * 前三，中三，后三、组三注数
     * @param WwNos
     * @return
     */
    public int setFrtTreGroup3(List<String> WwNos) {
        int ZS = 1;
        if(WwNos.size() <2) {
            return 0;
        }
        ZS = PlNum(WwNos.size(),2);
        return ZS;
    }

    /**
     * 前三，中三，后三，组六注数
     * @param WwNos
     * @return
     */
    public int setFrtTreGroup6(List<String> WwNos) {
        int ZS = 1;

        if(WwNos.size() <3) {
            return 0;
        }

        ZS = ZHnum(WwNos.size(),3);
        return ZS;
    }



    public int[] GroupResult = {1,2,2,4,5,6,8,10,11,13,14,14,15,15,14,14,13,11,10,8,6,5,4,2,2,1};
    public int[] ZXResult = {1,3,6,10,15,21,28,36,45,55,63,69,73,75,75,73,69,63,55,45,36,28,21,15,10,6,3,1};

    /**
     * 组选和值 三星
     * @return
     */
    public int GroupAndResult(List<String> GroupNos) {
        int ZS = 0;
        for (int i=0;i<GroupNos.size();i++) {
            ZS += GroupResult[Integer.valueOf(GroupNos.get(i))-1];
        }
        return ZS;
    }

    /**
     * 直选和值 三星
     * @return
     */
    public int ZXAndResult(List<String> DirectNo) {
        int ZS =0;
        List<Integer> Nos = new ArrayList<Integer>();
        for (int i =0;i<DirectNo.size();i++) {
            ZS += ZXResult[Integer.valueOf(DirectNo.get(i))];
        }
        return ZS;
    }

    /**
     * 二星直选注数
     * @param WwNos
     * @param QwNos
     * @return
     */
    public int setSecZX(List<String> WwNos,List<String> QwNos) {
        int ZS = 1;
        if(WwNos.size() <1) {
            return 0;
        }
        if (QwNos.size() <1){
            return 0;
        }

        ZS = WwNos.size()*QwNos.size();

        return ZS;
    }
    /**
     * 二星组选注数
     * @param WwNos
     * @return
     */
    public int setSecGroup(List<String> WwNos) {
        int ZS = 1;
        LottoryBalls balls = new LottoryBalls();
        if(WwNos.size() <2) {
            return 0;
        }
        ZS = balls.ZHnum(WwNos.size(),2);

        return ZS;
    }
    /**
     * 二星拖胆注数
     * @param WwNos
     * @param QwNos
     * @return
     */
    public int setSecTD(List<String> WwNos,List<String> QwNos) {
        int ZS = 1;
        if(WwNos.size()<1) {
            return 0;
        }
        if(QwNos.size()<1) {
            return 0;
        }

        ZS = QwNos.size();

        return ZS;
    }
    /**
     * 一星定位胆
     * @return
     */
    public int FrtStarDW(List<String> WwNos,List<String> QwNos,List<String> BwNos,List<String> SwNos,List<String> GwNos) {
        int ZS = 0;
        ZS = WwNos.size()+QwNos.size()+BwNos.size()+SwNos.size()+GwNos.size();
        return ZS;
    }
    /**
     * 一星直选注数
     * @return
     */
    public int FrtZX(List<String> WwNos) {
        return WwNos.size();
    }
    /**
     * 不定式一个号码注数
     * @param WwNos
     * @return
     */
    public  int BdsOneBall(List<String> WwNos) {
        int ZS = 0;
        ZS = WwNos.size();
        return ZS;
    }
    /**
     * 不定式两个号码注数
     * @param WwNos
     * @return
     */
    public int BdsTwoBalls(List<String> WwNos) {
        int ZS = 0;
        LottoryBalls balls = new LottoryBalls();
        ZS = balls.ZHnum(WwNos.size(),2);
        return ZS;
    }
    /**
     * 不定式三个号码注数
     * @param WwNos
     * @return
     */
    public int BdsTreBalls(List<String> WwNos) {
        int ZS = 0;
        LottoryBalls balls = new LottoryBalls();
        ZS = balls.ZHnum(WwNos.size(),3);
        return ZS;
    }

    /**
     * 任选三注数
     * @return
     */
    public int AnyTre(List<String> WwNos,List<String> QwNos,List<String> BwNos,List<String> SwNos,List<String> GwNos) {
        int i = WwNos.size();
        int j = QwNos.size();
        int k = BwNos.size();
        int m = SwNos.size();
        int n = GwNos.size();

        return k*m*n+j*m*n+j*k*n+j*k*m+i*m*n +i*k*n+i*k*m+i*j*n+i*j*m+i*j*k;
    }

    /**
     * 任选二注数
     * @return
     */
    public int AnyTwo(List<String> WwNos, List<String> QwNos, List<String> BwNos, List<String> SwNos, List<String> GwNos) {
        int i = WwNos.size();
        int j = QwNos.size();
        int k = BwNos.size();
        int m = SwNos.size();
        int n = GwNos.size();
        return i*j+i*k+i*m+i*n+j*k+j*m+j*n+k*m+k*n+m*n;
    }

    /**
     * 趣味注数
     * @return
     */
    public int QWBalls(List<String> WwNos) {
        return WwNos.size();
    }

}
