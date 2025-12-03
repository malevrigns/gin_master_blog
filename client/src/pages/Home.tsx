import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { ArrowRight, Github, Twitter, Linkedin, Instagram, Terminal, Cpu, Zap } from "lucide-react";
import { useState, useEffect } from "react";

export default function Home() {
  const [scrollY, setScrollY] = useState(0);

  useEffect(() => {
    const handleScroll = () => setScrollY(window.scrollY);
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  return (
    <div className="min-h-screen overflow-x-hidden bg-background text-foreground selection:bg-primary selection:text-black font-mono">
      {/* Dynamic Background */}
      <div className="fixed inset-0 -z-10 overflow-hidden pointer-events-none">
        <div 
          className="absolute top-0 left-0 w-full h-full opacity-60"
          style={{
            backgroundImage: 'url(/images/cyber-bg.png)',
            backgroundSize: 'cover',
            backgroundPosition: 'center',
            transform: `translateY(${scrollY * 0.1}px) scale(${1 + scrollY * 0.0002})`,
            filter: 'contrast(1.2) brightness(0.8)',
          }}
        />
        <div className="absolute inset-0 bg-background/80 mix-blend-multiply" />
        <div className="absolute inset-0 bg-[url('/images/cyber-shape-2.png')] opacity-10 mix-blend-screen animate-pulse" />
      </div>

      {/* HUD Overlay Elements - Moved outside background container */}
      <div className="fixed inset-0 z-40 pointer-events-none">
        {/* Top Left Corner */}
        <div className="absolute top-32 left-6 md:left-12 flex flex-col gap-2 opacity-70">
          <div className="flex items-center gap-2 text-[10px] text-primary font-mono tracking-widest">
            <div className="w-2 h-2 bg-primary animate-ping" />
            SYS.ROOT.ACCESS
          </div>
          <div className="h-px w-32 bg-gradient-to-r from-primary to-transparent" />
          <div className="text-[10px] text-muted-foreground font-mono">
            COORD: {(34 + scrollY * 0.001).toFixed(4)}N, {(118 - scrollY * 0.001).toFixed(4)}E
          </div>
        </div>

        {/* Top Right Corner */}
        <div className="absolute top-32 right-6 md:right-12 flex flex-col items-end gap-2 opacity-70">
          <div className="flex items-center gap-1">
            {[1,2,3,4,5].map(i => (
              <div key={i} className="w-1 h-4 bg-secondary/50" style={{ opacity: Math.random() }} />
            ))}
          </div>
          <div className="text-[10px] text-secondary font-mono tracking-widest">
            MEM: 64TB // NET: SECURE
          </div>
        </div>

        {/* Vertical Lines */}
        <div className="absolute top-0 left-8 md:left-16 w-px h-screen bg-white/5" />
        <div className="absolute top-0 right-8 md:right-16 w-px h-screen bg-white/5" />
      </div>

      {/* Navigation */}
      <nav className="fixed top-0 left-0 right-0 z-50 px-6 py-6 border-b border-primary/20 bg-background/80 backdrop-blur-md">
        <div className="container mx-auto flex items-center justify-between">
          <a href="/" className="text-2xl font-bold tracking-tighter hover:text-primary transition-colors flex items-center gap-2">
            <Terminal className="w-6 h-6 text-primary" />
            <span className="glitch" data-text="Gin.">Gin.</span>
          </a>
          <div className="hidden md:flex items-center gap-8">
            {[
              { label: '作品', href: '#work' },
              { label: '关于', href: '#about' },
              { label: '博客', href: '#blog' },
              { label: '联系', href: '#contact' }
            ].map((item) => (
              <a key={item.label} href={item.href} className="text-sm font-bold uppercase tracking-widest hover:text-primary transition-colors relative group">
                <span className="text-primary/50 mr-1">&lt;</span>
                {item.label}
                <span className="text-primary/50 ml-1">/&gt;</span>
              </a>
            ))}
          </div>
          <Button variant="outline" className="btn-neon rounded-none border-primary text-primary hover:bg-primary hover:text-black">
            初始化连接
          </Button>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="relative min-h-screen flex items-center pt-20 overflow-hidden">
        <div className="container mx-auto grid lg:grid-cols-2 gap-12 items-center relative z-10">
          <div className="space-y-8">
            <div className="inline-flex items-center gap-2 px-4 py-1 border border-secondary/50 bg-secondary/10 text-xs font-bold text-secondary uppercase tracking-[0.2em]">
              <span className="w-2 h-2 bg-secondary animate-pulse" />
              系统在线 :: 接受委托
            </div>
            
            <h1 className="text-5xl md:text-7xl font-bold leading-[0.9] tracking-tighter">
              <span className="block text-white mb-2">数字世界的</span>
              <span className="glitch text-primary block" data-text="赛博炼金术士">赛博炼金术士</span>
            </h1>
            
            <p className="text-lg md:text-xl text-muted-foreground max-w-lg leading-relaxed border-l-2 border-primary/50 pl-6">
              在代码与霓虹的交织中重构现实。打造融合高科技美学与极致交互的数字体验。
            </p>
            
            <div className="flex flex-wrap gap-4 pt-4">
              <a href="#work">
                <Button size="lg" className="btn-neon h-14 px-8 text-lg rounded-none">
                  加载项目模块 <Cpu className="ml-2 h-5 w-5" />
                </Button>
              </a>
              <a href="#blog">
                <Button size="lg" variant="outline" className="btn-neon-cyan h-14 px-8 text-lg rounded-none border-secondary text-secondary hover:bg-secondary hover:text-black">
                  访问数据日志
                </Button>
              </a>
            </div>

            <div className="flex items-center gap-6 pt-8 text-muted-foreground">
              {[Github, Twitter, Linkedin, Instagram].map((Icon, i) => (
                <a key={i} href="#" className="hover:text-primary hover:scale-110 transition-all duration-300 p-2 border border-transparent hover:border-primary/50 hover:bg-primary/10">
                  <Icon className="h-5 w-5" />
                </a>
              ))}
            </div>
          </div>

          <div className="relative hidden lg:flex h-[600px] w-full items-center justify-center">
            <img 
              src="/images/cyber-shape-1.png" 
              alt="Cyber Shape" 
              className="max-w-full max-h-[90%] object-contain animate-float drop-shadow-[0_0_50px_rgba(255,0,255,0.3)] z-20"
              style={{ animationDuration: '4s' }}
            />
            <div className="absolute inset-0 bg-gradient-to-t from-background via-transparent to-transparent z-10 pointer-events-none" />
          </div>
        </div>
        
        {/* Decorative Elements */}
        <div className="absolute bottom-0 left-0 w-full h-px bg-gradient-to-r from-transparent via-primary to-transparent opacity-50" />
        <div className="absolute top-1/4 right-0 w-64 h-64 bg-primary/20 blur-[100px] rounded-full pointer-events-none" />
        <div className="absolute bottom-1/4 left-0 w-64 h-64 bg-secondary/20 blur-[100px] rounded-full pointer-events-none" />
      </section>

      {/* Featured Work Section */}
      <section id="work" className="py-32 relative border-t border-white/5">
        <div className="container mx-auto">
          <div className="flex flex-col md:flex-row justify-between items-end mb-16 gap-6">
            <div>
              <h2 className="text-4xl md:text-5xl font-bold mb-4 flex items-center gap-4">
                <span className="text-primary text-2xl">01.</span> 精选作品
              </h2>
              <p className="text-muted-foreground text-lg max-w-md font-mono">/// 执行项目检索协议...</p>
            </div>
            <Button variant="ghost" className="group text-lg hover:text-primary hover:bg-transparent">
              查看所有数据 <ArrowRight className="ml-2 h-5 w-5 transition-transform group-hover:translate-x-1" />
            </Button>
          </div>

          <div className="grid md:grid-cols-2 gap-8">
            {[
              {
                id: 1,
                title: "Neural Nexus // 神经网络仪表盘",
                desc: "基于 AI 的实时数据可视化平台，监控全球神经网络节点的运行状态与数据流向。",
                tags: ["React", "D3.js", "AI Integration"],
                image: "/images/project-1.png"
              },
              {
                id: 2,
                title: "Neon City VR // 沉浸式漫游",
                desc: "第一人称视角的赛博朋克城市漫游体验，集成全息菜单与交互式 HUD 系统。",
                tags: ["WebXR", "Three.js", "WebGL"],
                image: "/images/project-2.png"
              },
              {
                id: 3,
                title: "DeFi Matrix // 去中心化交易终端",
                desc: "下一代加密货币交易界面，提供毫秒级数据更新与深色模式下的极致视觉体验。",
                tags: ["Next.js", "WebSocket", "Blockchain"],
                image: "/images/project-3.png"
              },
              {
                id: 4,
                title: "Glitch Gallery // 生成艺术画廊",
                desc: "探索数字故障美学的在线画廊，每一幅作品都由算法实时生成，独一无二。",
                tags: ["Canvas API", "Generative Art", "Algorithmic"],
                image: "/images/project-4.png"
              }
            ].map((item) => (
              <div key={item.id} className={`group relative ${item.id % 2 === 0 ? 'md:translate-y-24' : ''}`}>
                <Card className="cyber-panel rounded-none border-0 bg-black/50 overflow-visible">
                  <CardContent className="p-0 relative aspect-[4/3] overflow-hidden">
                    <img 
                      src={item.image} 
                      alt={item.title}
                      className="w-full h-full object-cover transition-transform duration-700 group-hover:scale-110 group-hover:filter group-hover:contrast-125"
                    />
                    <div className="absolute inset-0 bg-gradient-to-t from-background via-background/20 to-transparent opacity-60" />
                    
                    {/* Glitch Overlay on Hover */}
                    <div className="absolute inset-0 bg-[url('/images/cyber-shape-2.png')] opacity-0 group-hover:opacity-20 transition-opacity duration-100 mix-blend-overlay" />
                    
                    <div className="absolute bottom-0 left-0 right-0 p-6 bg-black/90 border-t border-primary/30 backdrop-blur-md translate-y-0 transition-transform duration-300">
                      <div className="flex justify-between items-center mb-2">
                        <div className="text-xs text-primary font-mono">PROJECT_0{item.id} // DEPLOYED</div>
                        <div className="h-1 w-12 bg-secondary/30 overflow-hidden">
                          <div className="h-full w-1/2 bg-secondary animate-pulse" />
                        </div>
                      </div>
                      <h3 className="text-xl font-bold mb-2 group-hover:text-primary transition-colors truncate">
                        {item.title}
                      </h3>
                      <p className="text-muted-foreground text-sm mb-4 line-clamp-2">
                        {item.desc}
                      </p>
                      <div className="flex gap-2 flex-wrap">
                        {item.tags.map((tag) => (
                          <span key={tag} className="text-[10px] px-2 py-1 border border-white/10 text-white/60 bg-white/5 hover:border-primary/50 hover:text-primary transition-colors cursor-default">
                            {tag}
                          </span>
                        ))}
                      </div>
                    </div>
                    
                    {/* Corner Accents */}
                    <div className="absolute top-0 left-0 w-4 h-4 border-t-2 border-l-2 border-primary transition-all duration-300 group-hover:w-8 group-hover:h-8" />
                    <div className="absolute top-0 right-0 w-4 h-4 border-t-2 border-r-2 border-primary transition-all duration-300 group-hover:w-8 group-hover:h-8" />
                    <div className="absolute bottom-0 left-0 w-4 h-4 border-b-2 border-l-2 border-primary transition-all duration-300 group-hover:w-8 group-hover:h-8" />
                    <div className="absolute bottom-0 right-0 w-4 h-4 border-b-2 border-r-2 border-primary transition-all duration-300 group-hover:w-8 group-hover:h-8" />
                  </CardContent>
                </Card>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* About Section */}
      <section id="about" className="py-32 relative overflow-hidden bg-black/50">
        <div className="absolute inset-0 bg-[url('/images/cyber-shape-2.png')] opacity-5" />
        
        <div className="container mx-auto relative z-10">
          <div className="max-w-4xl mx-auto text-center">
            <span className="text-secondary font-bold tracking-[0.3em] uppercase mb-6 block">关于系统核心</span>
            <h2 className="text-3xl md:text-5xl font-bold leading-tight mb-12">
              <span className="glitch" data-text="优秀的设计">优秀的设计</span> 不仅是视觉的呈现，<br/>更是 <span className="text-primary">灵魂的共鸣</span>。
            </h2>
            
            <div className="grid md:grid-cols-2 gap-12 text-lg text-muted-foreground leading-relaxed text-left border border-white/10 p-8 bg-black/40 backdrop-blur-sm relative">
              <div className="absolute -top-1 -left-1 w-3 h-3 bg-primary" />
              <div className="absolute -bottom-1 -right-1 w-3 h-3 bg-secondary" />
              
              <p>
                <span className="text-primary font-bold">&gt;</span> 拥有超过5年的数字产品设计和前端开发经验，我专注于创造令人印象深刻的沉浸式网络体验。
              </p>
              <p>
                <span className="text-secondary font-bold">&gt;</span> 我的方法结合了技术精确性与艺术直觉，确保每一个像素都有其存在的意义，每一次交互都在讲述一个故事。
              </p>
            </div>

            <div className="mt-16 grid grid-cols-2 md:grid-cols-4 gap-4">
              {[
                { label: '运行时间', value: '5Y+' },
                { label: '已编译项目', value: '50+' },
                { label: '节点连接', value: '30+' },
                { label: '获得成就', value: '12' },
              ].map((stat, i) => (
                <div key={i} className="border border-white/10 p-6 hover:border-primary/50 hover:bg-primary/5 transition-all group">
                  <div className="text-3xl font-bold text-white mb-2 group-hover:text-primary group-hover:scale-110 transition-all duration-300">{stat.value}</div>
                  <div className="text-xs text-muted-foreground uppercase tracking-widest">{stat.label}</div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* Blog Section */}
      <section id="blog" className="py-32 relative">
        <div className="container mx-auto">
          <h2 className="text-4xl md:text-5xl font-bold mb-16 text-center flex items-center justify-center gap-4">
            <span className="text-secondary text-2xl">02.</span> 数据日志
          </h2>
          
          <div className="grid md:grid-cols-3 gap-8">
            {[
              { title: "Web 动画的未来趋势", date: "2023-10-24", tag: "设计" },
              { title: "精通赛博朋克美学", date: "2023-11-12", tag: "教程" },
              { title: "为什么排版至关重要", date: "2023-12-05", tag: "观点" }
            ].map((post, i) => (
              <a key={i} href="#" className="group block">
                <article className="cyber-panel h-full p-6 hover:border-secondary/50 transition-all duration-300 hover:-translate-y-2">
                  <div className="mb-6 aspect-video bg-black border border-white/10 overflow-hidden relative">
                    <div className="absolute inset-0 bg-gradient-to-br from-primary/20 to-secondary/20 group-hover:opacity-100 opacity-50 transition-opacity" />
                    <Zap className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-white/20 w-12 h-12 group-hover:text-secondary group-hover:scale-125 transition-all duration-500" />
                  </div>
                  <div className="flex items-center gap-4 text-xs font-bold uppercase tracking-wider text-muted-foreground mb-4">
                    <span className="text-secondary px-2 py-1 border border-secondary/30 bg-secondary/10">{post.tag}</span>
                    <span className="font-mono">{post.date}</span>
                  </div>
                  <h3 className="text-xl font-bold mb-4 group-hover:text-primary transition-colors">{post.title}</h3>
                  <p className="text-sm text-muted-foreground font-mono">
                    正在解密数据包... 点击以访问完整内容。
                  </p>
                </article>
              </a>
            ))}
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-20 border-t border-primary/20 bg-black relative overflow-hidden">
        <div className="container mx-auto text-center relative z-10">
          <h2 className="text-4xl md:text-6xl font-bold mb-8">
            准备好 <span className="text-primary">接入</span> 了吗？
          </h2>
          <Button size="lg" className="btn-neon h-16 px-12 text-xl rounded-none mb-16">
            启动项目协议
          </Button>
          
          <div className="flex justify-between items-end pt-16 border-t border-white/10">
            <div className="text-left">
              <div className="text-2xl font-bold mb-2 flex items-center gap-2">
                <Terminal className="w-5 h-5 text-primary" /> Gin.
              </div>
              <p className="text-muted-foreground text-xs font-mono">© 2024 Gin Master Blog. 系统正常运行。</p>
            </div>
            <div className="flex gap-6 text-xs font-mono uppercase tracking-widest">
              <a href="#" className="text-muted-foreground hover:text-primary transition-colors">[ 隐私协议 ]</a>
              <a href="#" className="text-muted-foreground hover:text-primary transition-colors">[ 服务条款 ]</a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}
