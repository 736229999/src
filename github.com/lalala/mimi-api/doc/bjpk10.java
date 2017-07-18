package Utils;

import java.util.List;

/**
 * Created by zl on 2017/4/30.
 */

public class bjpk10 {

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
     * 猜冠军注数/位置一/拖拉机/猜奇偶/猜大小/和值
     * @param oneLineNos
     * @return
     */
    public int oneLine(List<String> oneLineNos) {
        int ZS=0;
        ZS = oneLineNos.size();
        return ZS;
    }
    /**
     * 猜冠亚军注数/精确前二
     * @param oneLineNos
     * @return
     */
    public int twoLine(List<String> oneLineNos,List<String> twoLineNos) {
        int ZS=0;
        int com = 0;
        for (int i =0;i<oneLineNos.size();i++) {
            for (int j=0;j<twoLineNos.size();j++) {
                if(oneLineNos.get(i).equals(twoLineNos.get(j))) {
                    com += 1;
                }
            }
        }
        ZS = oneLineNos.size()*twoLineNos.size()-com;
        return ZS;
    }

    /**
     * 精确组选2/位置二
     * @param oneLineNos
     * @return
     */
    public int Group2(List<String> oneLineNos) {
        int Zs = 0;
        Zs = ZHnum(oneLineNos.size(),2);
        return Zs;
    }
    /**
     * 精确组选3
     * @param oneLineNos
     * @return
     */
    public int Group3(List<String> oneLineNos) {
        int Zs = 0;
        Zs = ZHnum(oneLineNos.size(),3);
        return Zs;
    }
    /**
     * 精确组选4
     * @param oneLineNos
     * @return
     */
    public int Group4(List<String> oneLineNos) {
        int Zs = 0;

        Zs = ZHnum(oneLineNos.size(),4);
        return Zs;
    }
}
