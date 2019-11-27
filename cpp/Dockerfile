FROM amazonlinux

WORKDIR /app
COPY ./CMakeLists.txt ./main.cpp /app/

RUN yum update -y && yum install -y \
  gcc-c++ \
  cmake3 \
  git \
  make \
  zip \
  libcurl-devel \
  tar \
  wget \
  # Install aws-lambda-cpp
  && mkdir -p /app/build \
  && cd /app \
  && git clone https://github.com/awslabs/aws-lambda-cpp.git \
  && cd /app/aws-lambda-cpp \
  && mkdir build \
  && cd build \
  && cmake3 .. -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=~/lambda-install \
  && make \
  && make install \
  # Cleanup
  && yum clean all \
  && rm -rf /var/cache/yum

RUN cd /app/build \
  && cmake3 .. -DCMAKE_BUILD_TYPE=Debug -DCMAKE_INSTALL_PREFIX=~/lambda-install \
  && make \
  && make aws-lambda-package-cpp
