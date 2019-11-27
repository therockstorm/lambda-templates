#include <aws/lambda-runtime/runtime.h>

using namespace aws::lambda_runtime;

static invocation_response handler(invocation_request const& req)
{
    return invocation_response::success("response body", "application/json");
}

int main()
{
    run_handler(handler);
    return 0;
}
