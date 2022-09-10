/**
 * Copyright 2022 Ryan McGuinness
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 **/
package example.utils;

import com.google.protobuf.Any;
import com.google.protobuf.ByteString;
import com.google.protobuf.Timestamp;
import java.io.UnsupportedEncodingException;
import java.time.Instant;

public class PBUtils {
    private PBUtils() {

    }

    public static Timestamp getCurrentTime() {
        final Instant now = Instant.now();
        return Timestamp.newBuilder()
            .setSeconds(now.getEpochSecond())
            .setNanos(now.getNano()).build();
      }

      public static Any stringToAny(final String value) throws UnsupportedEncodingException {
        return Any.newBuilder()
            .setValue(ByteString.copyFrom(value, "UTF-8"))
            .build();
      }
}
