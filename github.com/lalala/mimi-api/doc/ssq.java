package Utils;

import java.util.List;

/**
 * Created by zl on 2017/4/30.
 */

public class Ssq {
    public int ssq(List<String> red, List<String> redTM, List<String> blue) {

        return ZHnum(redTM.size(),6-red.size()) * blue.size();
    }
}
