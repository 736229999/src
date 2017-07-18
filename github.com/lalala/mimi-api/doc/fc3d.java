package Utils;

import java.util.List;

/**
 * Created by zl on 2017/4/30.
 */

public class fC3D {

    //=====================================直选===============================
    public int FcZx(List<String> oneLine, List<String> twoLine, List<String> threeLine) {

        return oneLine.size()*twoLine.size()*threeLine.size();

    }

  //  ======================================组三==================================
    public int FzZs(List<String> oneLine) {

        return  PlNum(oneLine.size(),2);
    }

  //  ==================================================组六=========================
    public int FzZl(List<String> oneLine) {

        return  ZHnum(oneLine.size(),3);
    }

}
