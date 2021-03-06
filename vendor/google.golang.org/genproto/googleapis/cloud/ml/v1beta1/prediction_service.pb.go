// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/ml/v1beta1/prediction_service.proto

package ml

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_api3 "google.golang.org/genproto/googleapis/api/httpbody"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request for predictions to be issued against a trained model.
//
// The body of the request is a single JSON object with a single top-level
// field:
//
// <dl>
//   <dt>instances</dt>
//   <dd>A JSON array containing values representing the instances to use for
//       prediction.</dd>
// </dl>
//
// The structure of each element of the instances list is determined by your
// model's input definition. Instances can include named inputs or can contain
// only unlabeled values.
//
// Not all data includes named inputs. Some instances will be simple
// JSON values (boolean, number, or string). However, instances are often lists
// of simple values, or complex nested lists. Here are some examples of request
// bodies:
//
// CSV data with each row encoded as a string value:
// <pre>
// {"instances": ["1.0,true,\\"x\\"", "-2.0,false,\\"y\\""]}
// </pre>
// Plain text:
// <pre>
// {"instances": ["the quick brown fox", "la bruja le dio"]}
// </pre>
// Sentences encoded as lists of words (vectors of strings):
// <pre>
// {
//   "instances": [
//     ["the","quick","brown"],
//     ["la","bruja","le"],
//     ...
//   ]
// }
// </pre>
// Floating point scalar values:
// <pre>
// {"instances": [0.0, 1.1, 2.2]}
// </pre>
// Vectors of integers:
// <pre>
// {
//   "instances": [
//     [0, 1, 2],
//     [3, 4, 5],
//     ...
//   ]
// }
// </pre>
// Tensors (in this case, two-dimensional tensors):
// <pre>
// {
//   "instances": [
//     [
//       [0, 1, 2],
//       [3, 4, 5]
//     ],
//     ...
//   ]
// }
// </pre>
// Images can be represented different ways. In this encoding scheme the first
// two dimensions represent the rows and columns of the image, and the third
// contains lists (vectors) of the R, G, and B values for each pixel.
// <pre>
// {
//   "instances": [
//     [
//       [
//         [138, 30, 66],
//         [130, 20, 56],
//         ...
//       ],
//       [
//         [126, 38, 61],
//         [122, 24, 57],
//         ...
//       ],
//       ...
//     ],
//     ...
//   ]
// }
// </pre>
// JSON strings must be encoded as UTF-8. To send binary data, you must
// base64-encode the data and mark it as binary. To mark a JSON string
// as binary, replace it with a JSON object with a single attribute named `b64`:
// <pre>{"b64": "..."} </pre>
// For example:
//
// Two Serialized tf.Examples (fake data, for illustrative purposes only):
// <pre>
// {"instances": [{"b64": "X5ad6u"}, {"b64": "IA9j4nx"}]}
// </pre>
// Two JPEG image byte strings (fake data, for illustrative purposes only):
// <pre>
// {"instances": [{"b64": "ASa8asdf"}, {"b64": "JLK7ljk3"}]}
// </pre>
// If your data includes named references, format each instance as a JSON object
// with the named references as the keys:
//
// JSON input data to be preprocessed:
// <pre>
// {
//   "instances": [
//     {
//       "a": 1.0,
//       "b": true,
//       "c": "x"
//     },
//     {
//       "a": -2.0,
//       "b": false,
//       "c": "y"
//     }
//   ]
// }
// </pre>
// Some models have an underlying TensorFlow graph that accepts multiple input
// tensors. In this case, you should use the names of JSON name/value pairs to
// identify the input tensors, as shown in the following exmaples:
//
// For a graph with input tensor aliases "tag" (string) and "image"
// (base64-encoded string):
// <pre>
// {
//   "instances": [
//     {
//       "tag": "beach",
//       "image": {"b64": "ASa8asdf"}
//     },
//     {
//       "tag": "car",
//       "image": {"b64": "JLK7ljk3"}
//     }
//   ]
// }
// </pre>
// For a graph with input tensor aliases "tag" (string) and "image"
// (3-dimensional array of 8-bit ints):
// <pre>
// {
//   "instances": [
//     {
//       "tag": "beach",
//       "image": [
//         [
//           [138, 30, 66],
//           [130, 20, 56],
//           ...
//         ],
//         [
//           [126, 38, 61],
//           [122, 24, 57],
//           ...
//         ],
//         ...
//       ]
//     },
//     {
//       "tag": "car",
//       "image": [
//         [
//           [255, 0, 102],
//           [255, 0, 97],
//           ...
//         ],
//         [
//           [254, 1, 101],
//           [254, 2, 93],
//           ...
//         ],
//         ...
//       ]
//     },
//     ...
//   ]
// }
// </pre>
// If the call is successful, the response body will contain one prediction
// entry per instance in the request body. If prediction fails for any
// instance, the response body will contain no predictions and will contian
// a single error entry instead.
type PredictRequest struct {
	// Required. The resource name of a model or a version.
	//
	// Authorization: requires `Viewer` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	//
	// Required. The prediction request body.
	HttpBody *google_api3.HttpBody `protobuf:"bytes,2,opt,name=http_body,json=httpBody" json:"http_body,omitempty"`
}

func (m *PredictRequest) Reset()                    { *m = PredictRequest{} }
func (m *PredictRequest) String() string            { return proto.CompactTextString(m) }
func (*PredictRequest) ProtoMessage()               {}
func (*PredictRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *PredictRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PredictRequest) GetHttpBody() *google_api3.HttpBody {
	if m != nil {
		return m.HttpBody
	}
	return nil
}

func init() {
	proto.RegisterType((*PredictRequest)(nil), "google.cloud.ml.v1beta1.PredictRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OnlinePredictionService service

type OnlinePredictionServiceClient interface {
	// Performs prediction on the data in the request.
	//
	// **** REMOVE FROM GENERATED DOCUMENTATION
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*google_api3.HttpBody, error)
}

type onlinePredictionServiceClient struct {
	cc *grpc.ClientConn
}

func NewOnlinePredictionServiceClient(cc *grpc.ClientConn) OnlinePredictionServiceClient {
	return &onlinePredictionServiceClient{cc}
}

func (c *onlinePredictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*google_api3.HttpBody, error) {
	out := new(google_api3.HttpBody)
	err := grpc.Invoke(ctx, "/google.cloud.ml.v1beta1.OnlinePredictionService/Predict", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OnlinePredictionService service

type OnlinePredictionServiceServer interface {
	// Performs prediction on the data in the request.
	//
	// **** REMOVE FROM GENERATED DOCUMENTATION
	Predict(context.Context, *PredictRequest) (*google_api3.HttpBody, error)
}

func RegisterOnlinePredictionServiceServer(s *grpc.Server, srv OnlinePredictionServiceServer) {
	s.RegisterService(&_OnlinePredictionService_serviceDesc, srv)
}

func _OnlinePredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlinePredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.ml.v1beta1.OnlinePredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlinePredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OnlinePredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.ml.v1beta1.OnlinePredictionService",
	HandlerType: (*OnlinePredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _OnlinePredictionService_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/ml/v1beta1/prediction_service.proto",
}

func init() { proto.RegisterFile("google/cloud/ml/v1beta1/prediction_service.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4d, 0x4a, 0x03, 0x31,
	0x14, 0xc7, 0x49, 0x11, 0xb5, 0x11, 0x5c, 0x04, 0xb1, 0xb5, 0x08, 0x96, 0xba, 0xb0, 0x74, 0x91,
	0xd8, 0xba, 0xb2, 0xe2, 0xa6, 0x2b, 0x77, 0x0e, 0x75, 0x21, 0xb8, 0x29, 0xe9, 0x4c, 0x48, 0x23,
	0x99, 0xbc, 0x38, 0x93, 0x16, 0x8b, 0xb8, 0xf1, 0x0a, 0x3d, 0x9a, 0x57, 0xf0, 0x20, 0x92, 0x49,
	0x28, 0xca, 0xe8, 0xee, 0x31, 0x6f, 0x7e, 0xef, 0xff, 0x11, 0x7c, 0x29, 0x01, 0xa4, 0x16, 0x2c,
	0xd5, 0xb0, 0xcc, 0x58, 0xae, 0xd9, 0x6a, 0x38, 0x17, 0x8e, 0x0f, 0x99, 0x2d, 0x44, 0xa6, 0x52,
	0xa7, 0xc0, 0xcc, 0x4a, 0x51, 0xac, 0x54, 0x2a, 0xa8, 0x2d, 0xc0, 0x01, 0x69, 0x05, 0x82, 0x56,
	0x04, 0xcd, 0x35, 0x8d, 0x44, 0xe7, 0x34, 0x9e, 0xe2, 0x56, 0x31, 0x6e, 0x0c, 0x38, 0xee, 0xe9,
	0x32, 0x60, 0x9d, 0x93, 0x1f, 0xdb, 0x85, 0x73, 0x76, 0x0e, 0xd9, 0x3a, 0xac, 0x7a, 0x8f, 0xf8,
	0x30, 0x09, 0x6a, 0x53, 0xf1, 0xb2, 0x14, 0xa5, 0x23, 0x04, 0xef, 0x18, 0x9e, 0x8b, 0x36, 0xea,
	0xa2, 0x7e, 0x73, 0x5a, 0xcd, 0x64, 0x88, 0x9b, 0x9e, 0x9b, 0x79, 0xb0, 0xdd, 0xe8, 0xa2, 0xfe,
	0xc1, 0xe8, 0x88, 0x46, 0x2f, 0xdc, 0x2a, 0x7a, 0xe7, 0x9c, 0x9d, 0x40, 0xb6, 0x9e, 0xee, 0x2f,
	0xe2, 0x34, 0xda, 0x20, 0xdc, 0xba, 0x37, 0x5a, 0x19, 0x91, 0x6c, 0xd3, 0x3c, 0x84, 0x30, 0xe4,
	0x15, 0xef, 0xc5, 0x8f, 0xe4, 0x82, 0xfe, 0x13, 0x89, 0xfe, 0xb6, 0xd5, 0xf9, 0x53, 0xaf, 0x47,
	0x3f, 0x3e, 0xbf, 0x36, 0x8d, 0x7e, 0xef, 0x7c, 0xdb, 0xdd, 0x9b, 0x37, 0x7c, 0x6b, 0x0b, 0x78,
	0x16, 0xa9, 0x2b, 0xd9, 0x60, 0xf0, 0x3e, 0x8e, 0x75, 0x8e, 0xd1, 0x60, 0xb2, 0xc2, 0x67, 0x29,
	0xe4, 0x35, 0x4d, 0x7f, 0x33, 0x1e, 0x98, 0x1c, 0xd7, 0xfc, 0x26, 0xbe, 0xa9, 0x04, 0x3d, 0x5d,
	0x47, 0x4c, 0x82, 0xe6, 0x46, 0x52, 0x28, 0x24, 0x93, 0xc2, 0x54, 0x3d, 0xb2, 0xb0, 0xe2, 0x56,
	0x95, 0xb5, 0xe7, 0xbc, 0xc9, 0xf5, 0x7c, 0xb7, 0xfa, 0xeb, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0x00, 0x26, 0x25, 0x67, 0xf3, 0x01, 0x00, 0x00,
}
