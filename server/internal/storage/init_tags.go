package storage

import (
	"github.com/imdotdev/im.dev/server/internal/tags"
	"github.com/imdotdev/im.dev/server/pkg/models"
)

// \n\n\n\u003cWebsiteLink type=\"official\" url=\"https://reactjs.org/\" /\u003e
var tagsList = []*models.Tag{
	&models.Tag{Title: "Javascript", Name: "javascript", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1607082785538/EryuLRriM.png?w=400&h=400&fit=crop&crop=entropy&auto=compress", Md: "JavaScript is a cross-platform, object-oriented scripting language. It is a small and lightweight language. Inside a host environment (for example, a web browser), JavaScript can be connected to the objects of its environment to provide programmatic control over them."},
	&models.Tag{Title: "React", Name: "reactjs", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1513321478077/ByCWNxZMf.png?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "React is a JavaScript library for creating user interfaces by Facebook and Instagram. Many people choose to think of React as the V in MVC. React solves one problem: building large applications with data that changes over time\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://reactjs.org/\" /\u003e"},
	&models.Tag{Title: "Nodejs", Name: "nodejs", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1607322143407/LsbeyUL86.png?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "Node.js is an open-source, cross-platform, back-end, JavaScript runtime environment that executes JavaScript code outside a web browser.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://nodejs.org/en//\" /\u003e"},
	&models.Tag{Title: "Python", Name: "python", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1607321670861/lG265gIUp.png?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "Python is an interpreted, object-oriented, high-level programming language with dynamic semantics. Its high-level built-in data structures, combined with dynamic typing and dynamic binding, make it very attractive for Rapid Application Development, as well as for use as a scripting or glue language to connect existing components together.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://python.org/\" /\u003e"},
	&models.Tag{Title: "CSS", Name: "css", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1513316949083/By6UMkbfG.png?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "Cascading Style Sheets (CSS) is a stylesheet language used to describe the presentation of a document written in HTML. CSS describes how elements should be rendered on screen, on paper, in speech, or on other media"},
	&models.Tag{Title: "Java", Name: "java", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1534512378322/H1gM-pH4UQ.png?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "Java is a programming language and computing platform first released by Sun Microsystems in 1995. There are lots of applications and websites that will not work unless you have Java installed, and more are created every day. Java is fast, secure, and reliable. From laptops to datacenters, game consoles to scientific supercomputers, cell phones to the Internet, Java is everywhere.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://java.com/\" /\u003e"},
	&models.Tag{Title: "Go", Name: "go", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1534512687168/S1D40rVLm.png?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://golang.org/\" /\u003e"},
	&models.Tag{Title: "PHP", Name: "php", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1513177307594/rJ4Jba0-G.png?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "PHP is a server-side scripting language designed for web development but also used as a general-purpose programming language. Originally created by Rasmus Lerdorf in 1994, the PHP reference implementation is now produced by The PHP Group.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://php.net/\" /\u003e"},
	&models.Tag{Title: "React Native", Name: "react-native", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/rkij45wit50lfpkbte5q/1475235386.jpg?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "React Native enables you to build world-class application experiences on native platforms using a consistent developer experience based on JavaScript and React.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://reactnative.dev/\" /\u003e"},
	&models.Tag{Title: "Angular", Name: "angular", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1513839958045/HJ0LTRdMz.jpeg?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "Angular is a development platform for building mobile and desktop web applications.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://angular.io/\" /\u003e"},
	&models.Tag{Title: "Android", Name: "android", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/qbj34hxd8981nfdugyph/1450468271.png?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: "Android is a mobile operating system developed by Google, based on the Linux kernel and designed primarily for touchscreen mobile devices such as smartphones and tablets.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://android.com/\" /\u003e"},
	&models.Tag{Title: "Machine Learning", Name: "machine-learning", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1513321644252/Sk43El-fz.png?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: ""},
	&models.Tag{Title: "Developer", Name: "developer", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/v1554321431158/MqVqSHr8Q.jpeg?w=180&h=180&fit=crop&crop=entropy&auto=compress", Md: ""},
	&models.Tag{Title: "Docker", Name: "docker", Icon: "https://res.cloudinary.com/practicaldev/image/fetch/s--7L1txQC1--/c_limit,f_auto,fl_progressive,q_80,w_180/https://dev-to-uploads.s3.amazonaws.com/uploads/badge/badge_image/87/docker-badge.png", Md: "Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications, whether on laptops, data center VMs, or the cloud.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://reactjs.org/\" /\u003e"},
	&models.Tag{Title: "Kubernetes", Name: "kubernetes", Icon: "https://res.cloudinary.com/practicaldev/image/fetch/s--XXaJdQCT--/c_limit,f_auto,fl_progressive,q_80,w_180/https://dev-to-uploads.s3.amazonaws.com/uploads/badge/badge_image/88/kubernetes-badge.png", Md: "Kubernetes, also known as K8s, is an open-source system for automating deployment, scaling, and management of containerized applications.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://kubernetes.io/\" /\u003e"},
	&models.Tag{Title: "Rust", Name: "rust", Icon: "https://res.cloudinary.com/practicaldev/image/fetch/s--4RD0-X0v--/c_limit,f_auto,fl_progressive,q_80,w_180/https://dev-to-uploads.s3.amazonaws.com/uploads/badge/badge_image/35/rust-badge.png", Md: "Rust is a systems programming language that runs blazingly fast, prevents segfaults, and guarantees thread safety.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://rust-lang.org/\" /\u003e"},
	&models.Tag{Title: "Web Development", Name: "webdev", Icon: "https://i.postimg.cc/RVP8fWFz/web-dev-logo.png", Md: ""},
	&models.Tag{Title: "Testing", Name: "testing", Icon: "https://i.postimg.cc/0ygM0fzs/testing-image.png", Md: "Executing a program or application with the intent of finding the software bugs is known as Software Testing. It is normally done to detect the differences between given input and expected output."},
	&models.Tag{Title: "VueJS", Name: "vue", Icon: "https://res.cloudinary.com/practicaldev/image/fetch/s--KflrVC4o--/c_limit,f_auto,fl_progressive,q_80,w_180/https://dev-to-uploads.s3.amazonaws.com/uploads/badge/badge_image/27/vue-sticker.png", Md: "Vue is a progressive framework for building user interfaces. Unlike other monolithic frameworks, Vue is designed from the ground up to be incrementally adoptable.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://vuejs.org/\" /\u003e"},
	&models.Tag{Title: "Discuss", Name: "discuss", Icon: "https://i.postimg.cc/jScTwK9B/disscuss-image.png", Md: ""},
	&models.Tag{Title: "News", Name: "news", Icon: "https://i.postimg.cc/yxm5MyC2/news-image.png", Md: ""},
	&models.Tag{Title: "Typescript", Name: "typescript", Icon: "https://i.postimg.cc/PJTKwG0F/typescript-image.png", Md: "TypeScript is an open-source language that builds on JavaScript, one of the world's most used tools, by adding static type definitions. Types provide a way to describe the shape of an object, providing better documentation, and allowing TypeScript to validate that your code is working correctly.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://typescriptlang.org/\" /\u003e"},
	&models.Tag{Title: "Linux", Name: "linux", Icon: "https://cdn.hashnode.com/res/hashnode/image/upload/ogpsvoxw5kt8aksuiptj/1450641462.png?w=400&h=400&fit=crop&crop=entropy&auto=compress", Md: "Linux is a Unix-like and mostly POSIX-compliant computer OS assembled under the model of free and open-source software development and distribution.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://linux.org/\" /\u003e"},
	&models.Tag{Title: "Mysql", Name: "mysql", Icon: "https://i.postimg.cc/fyytFBdR/mysql-image.png", Md: "MySQL is an Open Source Relational SQL database management system.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://mysql.com/\" /\u003e"},
	&models.Tag{Title: "Flutter", Name: "flutter", Icon: "https://i.postimg.cc/dtp7D2bz/flutter-image.png", Md: "Flutter is an open-source UI software development kit created by Google. It is used to develop native applications for various platforms, and the web from a single codebase.\n\n\n\u003cWebsiteLink type=\"official\" url=\"https://flutter.dev/\" /\u003e"},
	&models.Tag{Title: "IOS", Name: "ios", Icon: "https://i.postimg.cc/z3ZFCsLg/ios-image.png", Md: "iOS (originally iPhone OS) is a mobile operating system created and developed by Apple Inc. and distributed exclusively for Apple hardware. It is the operating system that presently powers many of the company's mobile devices, including the iPhone, iPad, and iPod touch."},
	&models.Tag{Title: "CPP", Name: "cpp", Icon: "https://res.cloudinary.com/practicaldev/image/fetch/s--lRZA1qmc--/c_limit,f_auto,fl_progressive,q_80,w_180/https://dev-to-uploads.s3.amazonaws.com/uploads/badge/badge_image/95/cpp_logo.png", Md: "C++ is a general-purpose programming language which was developed by Bjarne Stroustrup at Bell Labs. It has imperative, object-oriented and generic programming features with support for low level memory manipulation."},
	&models.Tag{Title: "APM", Name: "apm", Icon: "https://i.postimg.cc/8cYwcyg7/apm-image.png", Md: "In the fields of information technology and systems management, application performance management (APM) is the monitoring and management of performance and availability of software applications. APM strives to detect and diagnose complex application performance problems to maintain an expected level of service."},
	&models.Tag{Title: "Algorithms", Name: "algorithms", Icon: "https://i.postimg.cc/yYBQtRpC/algorithms-image.png", Md: ""},
}

func initTags(u string) error {
	for _, tag := range tagsList {
		tag.Creator = u
		tags.SubmitTag(tag)
	}
	return nil
}
