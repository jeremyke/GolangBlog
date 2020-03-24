package message

//定义一个用户结构体
type User struct {
	UserName   string `json : "userName"` //为了·序列化跟反序列化成功，要给个tag
	UserId     int    `json : "userId"`
	UserPwd    string `json : "userPwd"`
	UserStatus int    `json : "userStatus"`
}

/*

public class TestException
{
    public TestException()
    {
    }
    boolean testEx() throws Exception
    {
        boolean ret = true;
        try
        {
            ret =testEx1();
        }
        catch (Exception e)
        {
            System.out.println("testEx,catch exception");
            ret =false;
            throwe;
        }
        finally
        {
            System.out.println("testEx,finally; return value=" + ret);
            returnret;
        }
    }
    boolean testEx1() throws Exception
    {
        boolean ret = true;
        try
        {
            ret =testEx2();
            if(!ret)
            {
                returnfalse;
            }
            System.out.println("testEx1,at the end of try");
            returnret;
        }
        catch (Exception e)
        {
            System.out.println("testEx1,catch exception");
            ret =false;
            throwe;
        }
        finally
        {
            System.out.println("testEx1,finally; return value=" + ret);
            returnret;
        }
    }
    boolean testEx2() throws Exception
    {
        boolean ret = true;
        try
        {
            int b =12;
            int c;
            for(int i = 2; i >= -2; i--)
            {
                c= b / i;
                System.out.println("i="+ i);
            }
            return true;
        }
        catch (Exception e)
        {
            System.out.println("testEx2,catch exception");
            ret =false;
            throwe;
        }
        finally
        {
            System.out.println("testEx2,finally; return value=" + ret);
            return ret;
        }
    }
    public static void main(String[] args)
    {
        TestException testException1 =new TestException();
        try
        {
            testException1.testEx();
        }
        catch (Exception e)
        {
            e.printStackTrace();
        }
    }
}

*/
